while getopts s: flag
do 
	case "${flag}" in
		s) opt=${OPTARG};;
	esac
done

case "$opt" in 
	"deploy")
		docker-compose --env-file=.env.dev up -d --build
		echo "Deployed"
		;;
	"undeploy")
		docker-compose --env-file=.env.dev stop
		echo "Undeployed"
		;;
esac
echo "Done"

