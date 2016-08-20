# env variables for go
cat >> /home/vagrant/.bashrc << EOF
export GOPATH=/vagrant
export GOBIN=/vagrant/bin
export PATH=$PATH:/vagrant/bin
export HOME=/home/vagrant
export ISGODEVENVIRONMENT="1"
EOF

export GOPATH=/vagrant
export GOBIN=/vagrant/bin
export PATH=$PATH:/vagrant/bin
export HOME=/home/vagrant
export ISGODEVENVIRONMENT="1"

git clone https://github.com/grameh/vimconf /home/vagrant/vimconf
ln -s  /home/vagrant/vimconf/.vimrc /home/vagrant/.vimrc
ln -s  /home/vagrant/vimconf/.vimrc.go /home/vagrant/.vimrc.go

echo "Installing vim plugins and binaries for vim-go :) this will take a bit"
vim  -c VundleUpdate -c quitall > /dev/null

echo "Compiling youcompleteme with go support"
cd /home/vagrant/.vim/bundle/YouCompleteMe
./install.py --gocode-completer

cd /vagrant/
echo "Getting project's dependencies"
go get
echo "building"
go build -o hello.exe

echo "done"
