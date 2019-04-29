<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: multipart/form-data");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");

$src = $_FILES['image']['tmp_name'];
$dest = "/home/teammkig/public_html/images/" . 
	$_POST["folder"] . 
	"/" . 
	$_POST["name"] . 
	"." . 
	pathinfo($_FILES['image']['name'], PATHINFO_EXTENSION);

if(move_uploaded_file($src, $dest)
) {
	print_r('moved');
} else {
    // Failed to move file
    print_r('failed');
}
