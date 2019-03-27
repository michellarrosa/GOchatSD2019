#!/bin/bash
# insira nome de USUÁRIO como abaixo, sem o ultimo slash
USERHOME="michel";
#nao modifique o codigo abaixo, comente/descomente, conforme instruido

# # # LINUX
#export GOROOT=/usr/lib/go-1.7; #debianico
#export GOROOT=/usr/lib/go-1.6; #outros

# export GOPATH=$USERHOME/Projects/GO;
    # export PATH=$PATH:$GOPATH/bin:$GOROOT/bin;
        # export GOBIN=$GOPATH/bin;
        
texto="export GOPATH=/home/$USERHOME/Projects/GO\nexport PATH=\$PATH:\$GOPATH/bin:$GOROOT/bin\nexport GOBIN=\$GOPATH/bin\n";
#only debian
echo $texto;
# # # Ubuntu não é linux
#export texto="export GOPATH=$USERHOME/Projects/GO\nexport PATH=\$PATH:\$GOPATH/bin:\$GOPATH/bin\nexport GOBIN=\$GOPATH/bin\n";
#echo -e $texto >> $USERHOME/.profile;
#source $USERHOME/.profile;

echo "# # # Ubuntu não é linux # # #";
