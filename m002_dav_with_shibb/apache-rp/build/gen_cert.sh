gen_cert() {
  fn=$1
  SVCN=$2
  SVDN="/C=FR/ST=Occitanie/L=Toulouse/O=otvl org/CN=${SVCN}"
  openssl req -new -x509 -days 3650 -nodes -out $fn.crt -keyout $fn.key -subj "$SVDN"
  openssl x509 -noout -text -in $fn.crt | egrep "Subject:|Issuer:|Serial Number:|Not After :"
  chmod 644 $fn.crt $fn.key
}

gen_cert pm02.otvl.org pm02.otvl.org
