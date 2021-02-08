APP_NAME=meonzi
CURRENT_PID=$(pgrep -f $APP_NAME)
if [ -z $CURRENT_PID ]
then
  echo "Running process is not exists"
else
  echo "kill -9 $CURRENT_PID"
  kill -15 $CURRENT_PID
  sleep 5
fi

echo "$APP_NAME start"
nohup /app/meonzi
