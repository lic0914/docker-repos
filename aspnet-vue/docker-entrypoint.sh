#!/bin/bash
set -e

# nginxProc=$(ps -ef |grep -w nginx|grep -v grep|wc -l)
# if [ "$nginxProc" -le 0 ] ;then
#     nginx -g daemon off;
# fi

update_config(){
   match=$1
   replace=$2
   echo "$match $replace"
   line=$(cat nginx.conf -n | grep -A5 server  | grep $match | awk '{print $1}')
   echo $line
   if [[ "$replace" =~ .*"/".* ]] ;then
      #echo "nice"
      sed  ''"$line"' s#'"$match"'[^\{]\+#'"$match"' '"$replace"' #' nginx.conf
   else
     # sed ''"$line"' s/'"$match"'*/'"$match"' '"$replace"'/' nginx.conf
     echo "hate"
   fi

}

args=(`echo "$*"`)

index=-1
ddl="";


for i in ${!args[@]}
do
echo "$i : ${args[i]}"
  arg=${args[i]}
  next=${args[i+1]}

  if [ "$arg" == "--dotnet" ] ;then
    ddl=$next
    continue
  fi

  if [ "$arg" == "--vue-prefix" ] ;then
    echo "vue-prefix $next"
    update_config "location" $next
  fi

  if [ "$arg" == "--vue-index-dir" ] ;then
    update_config "root" "$next;"
    continue
  fi

done

#dotnet $ddl


#--vue-prefix 修改nginx.conf
#--dotnet  
# nginx -V
