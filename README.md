Simple hello world docker image build with golang. 

Build an image 
$ docker build -t my-golang-app .

Run the app inside the image
$ docker run -it --rm --name my-running-app my-golang-app