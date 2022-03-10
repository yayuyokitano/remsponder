set -e

#set up cloud sql proxy
curl https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -o cloud_sql_proxy
chmod +x cloud_sql_proxy
./cloud_sql_proxy -instances=rem-970606:us-central1:rem=tcp:5432 &

#get the environment variables for tests, commented out for now but might be needed later.
#while read p; do
#  IFS=' = ' read -r -a envArray <<< "$p"
#  declare "${envArray[0]}=${envArray[1]}"
#  export ${envArray[0]}
#done < config.ini

go test .