# Placeholder image app

The application returns transparent images in the smallest ratio possible 
based on the requested size. This means the images can be used in HTML to 
hold a specific dimension and then CSS background images can be used to 
place responsive images, that only load images based on specified media 
queries.

By sending the smallest ratio possible, the file size can be small 
(extremely small in most cases) but the ratio is what is important when 
holding the place in HTML using background CSS images.

An example can be found here:
http://codepen.io/antimatter/pen/emBvPj

Transparent images are returned by specifying the width and height:

http://fpo-images.appspot.com/1920x1080

The application also returns placeholder images (i.e., single color images) 
to be used in wireframing or prototyping by specifying a color in the query 
string and/or requesting the fullsize image.

http://fpo-images.appspot.com/1920x1080?c=c0c0c0

http://fpo-images.appspot.com/1920x1080?c=c0c0c0&f=true

This was mostly an experiment with Golang and there are definitely some 
rough edges...

The app runs on Google App Engine