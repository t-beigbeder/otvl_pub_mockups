gen_cert() {
  fn=$1
  SVCN=$2
  SVDN="/C=FR/ST=Occitanie/L=Toulouse/O=otvl org/CN=${SVCN}"
  openssl req -x509 -newkey rsa:4096 -nodes -keyout $fn.key -out $fn.crt -days 3650 -subj "$SVDN" -addext "subjectAltName = DNS:$SVCN"
  openssl x509 -noout -text -in $fn.crt | egrep "Subject:|Issuer:|Serial Number:|Not After :"
  chmod 644 $fn.crt $fn.key
}

gen_cert pm02.srv.dxpydk pm02.srv.dxpydk
gen_cert pm02.otvl.org pm02.otvl.org
