set -e

env=""
while read p; do
  env+="$p=$(gcloud secrets versions access latest --secret "$p"),"
done < secretenv.txt

while read p; do
  IFS=' = ' read -r -a envArray <<< "$p"
  env+="${envArray[0]}=${envArray[1]},"
done < config.ini

gcloud functions deploy Respond --set-env-vars="${env%,}" --vpc-connector rem-connector --region=us-central1 --source . --trigger-topic responder --runtime go116