import logging
import os
import traceback
import argparse
import sys
import json

import tornado.ioloop
import tornado.web


logger = logging.getLogger(__name__)


def setup_env():
    logging.basicConfig(
        level=os.getenv('LOGGING', 'INFO'),
        format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')


class BaseHandler(tornado.web.RequestHandler):
    logger = logging.getLogger(__module__ + '.' + __qualname__)  # noqa
    server_config = {}
    site_config = {}
    j24bots_loader = None

    def _request_summary(self):
        if 'User-Agent' in self.request.headers:
            ua = self.request.headers['User-Agent']
        else:
            ua = "undefined_User-Agent"
        if 'X-Forwarded-For' in self.request.headers:
            xff = self.request.headers['X-Forwarded-For']
        else:
            xff = "undefined_X-Forwarded-For"

        s = "%s %s (%s) (%s)" % (
            self.request.method,
            self.request.uri,
            xff,
            ua
        )
        return s

    def prepare(self):
        if 'User-Agent' in self.request.headers:
            ua = self.request.headers['User-Agent']
        else:
            ua = "undefined_User-Agent"
        self.logger.debug(f"prepare: {ua} {self.request.method} {self.request.path}")
        if not self.request.path.endswith(".xml"):
            if "/html4/" not in self.request.path and self.request.path != "/api/html4":
                self.set_header("Content-Type", "application/json")
            else:
                self.set_header("Content-Type", "text/html")
        else:
            self.set_header("Content-Type", "text/xml; charset=utf-8")
        if "Origin" in self.request.headers and \
                "cors_mapping" in self.server_config and \
                self.request.headers["Origin"] in self.server_config["cors_mapping"]:
            origin_allowed = self.request.headers["Origin"]
            self.logger.debug(f"prepare CORS authorized for {origin_allowed}")
            self.set_header("Access-Control-Allow-Origin", origin_allowed)

    def _check_par(self, name, par):
        if not par:
            self._error(400, 'MissingParameter', 'Parameter {0} is missing in URL'.format(name))
            return par
        return True

    def _error(self, code, reason, message):
        self.set_status(code)
        self.finish({'reason': reason, 'message': message})


class VersionHandler(BaseHandler):
    logger = logging.getLogger(__module__ + '.' + __qualname__)  # noqa

    def initialize(self, **kwargs):
        super().initialize(**kwargs)

    def get(self):
        self.write(json.dumps({"version": "1.0"}, indent=2))
        return self.finish()


class AppServerMainBase:
    logger = logging.getLogger(__module__ + '.' + __qualname__)  # noqa

    def _arg_parser(self):
        raise NotImplementedError("_arg_parser")

    def __init__(self, name):
        self.name = name
        self.arg_parser = self._arg_parser()
        self.args = None

    def _do_run(self):
        raise NotImplementedError("_do_run")

    def run(self):
        self.args = self.arg_parser.parse_args()
        try:
            self.logger.info('run: start')
            result = self._do_run()
            self.logger.info('run: done')
            return result
        except Exception as e:
            traceback.print_exc()
            self.logger.error(
                'An unkonwn error occured, please contact the support - {0} {1}'.format(
                    type(e), e))
        self.logger.info('run: done')
        return False


def make_otvl_web_app():
    assets_directory = os.getenv("ASSETS_DIR", "/srv/assets")
    return tornado.web.Application([
        (r"/app/version/?", VersionHandler),
        (r"/app/assets/(.*)", tornado.web.StaticFileHandler, {"path": assets_directory}),
    ])


class OtvlWebServer(AppServerMainBase):
    logger = logging.getLogger(__module__ + "." + __qualname__)  # noqa

    @classmethod
    def _make_app(cls):
        return make_otvl_web_app()

    def _arg_parser(self):
        parser = argparse.ArgumentParser(description='OtvlWebServer')
        return parser

    def __init__(self, name):
        AppServerMainBase.__init__(self, name)

    def _do_run(self):
        self.logger.debug("_do_run OtvlWebServer")
        port = int(os.getenv("OW_PORT", "9090"))
        address = os.getenv("OW_ADDRESS", "")
        app = self._make_app()
        app.listen(port, address)
        tornado.ioloop.IOLoop.current().start()
        return True


setup_env()

if __name__ == "__main__":
    cmd_name = os.path.basename(sys.argv[0]).split('.')[0]
    res = OtvlWebServer(cmd_name).run()
    sys.exit(0 if res else -1)
