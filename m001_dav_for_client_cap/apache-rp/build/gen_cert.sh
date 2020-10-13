cd /srv/ssl

CACN="CA_otvl_web"
CADN="/C=FR/ST=Midi-Pyrenees/L=Toulouse/O=otvl org/CN=${CACN}"

if [ ! -f CA_otvl_web.csr ] ; then
  openssl req -new -newkey rsa:2048 -nodes -out CA_otvl_web.csr -keyout CA_otvl_web.key -subj "$CADN"
  openssl req -noout -text -in CA_otvl_web.csr
  openssl x509 -trustout -signkey CA_otvl_web.key -days 3650 -req -in CA_otvl_web.csr -out CA_otvl_web.crt

fi
openssl x509 -noout -text -in CA_otvl_web.crt | egrep "Subject:|Issuer:|Serial Number:|Not After :"

count=11

gen_cert() {
  count=`expr $count + 1`
  fn=$1
  SVCN=$2
  if [ ! -f $fn.crt ] ; then
    SVDN="/C=FR/ST=Midi-Pyrenees/L=Toulouse/O=otvl org/CN=${SVCN}"
    openssl req -new -newkey rsa:2048 -nodes -out $fn.req -keyout $fn.key -subj "$SVDN"
    echo $count > $fn.srl
    openssl x509 -CA CA_otvl_web.crt -CAkey CA_otvl_web.key -CAserial $fn.srl -days 3650 -req -in $fn.req -out $fn.crt
  fi
  openssl x509 -noout -text -in $fn.crt | egrep "Subject:|Issuer:|Serial Number:|Not After :"
}

gen_cert pm01.otvl.org pm01.otvl.org
