#!/bin/sh
CONF=/etc/config/qpkg.conf
QPKG_NAME="google-ddns-client"
QPKG_ROOT=`/sbin/getcfg $QPKG_NAME Install_Path -f ${CONF}`
APACHE_ROOT=`/sbin/getcfg SHARE_DEF defWeb -d Qweb -f /etc/config/def_share.info`
CLIENT_PID=/var/run/google-ddns-client.pid
export QNAP_QPKG=$QPKG_NAME

case "$1" in
  start)
    ENABLED=$(/sbin/getcfg $QPKG_NAME Enable -u -d FALSE -f $CONF)
    if [ "$ENABLED" != "TRUE" ]; then
        echo "$QPKG_NAME is disabled."
        exit 1
    fi

    # Kill any 'zombie' servers
    if [ -e $CLIENT_PID ]; then
        kill -9 $(cat $CLIENT_PID)
        rm -f $CLIENT_PID
    fi

    : ADD START ACTIONS HERE
    # Start the server
    $QPKG_ROOT/google-ddns-client -c $QPKG_ROOT/etc/config.ini > $QPKG_ROOT/google-ddns-client.log 2>&1 &

    # Capture the PID
    PID=$!
    echo Process ID $PID
    echo $PID > $CLIENT_PID 
    ;;

  stop)
    : ADD STOP ACTIONS HERE
    if [ -e $CLIENT_PID ]; then
        kill -9 $(cat $CLIENT_PID)
        rm -f $CLIENT_PID
    fi
    ;;

  restart)
    $0 stop
    $0 start
    ;;
  remove):
    ;;

  *)
    echo "Usage: $0 {start|stop|restart|remove}"
    exit 1
esac

exit 0
