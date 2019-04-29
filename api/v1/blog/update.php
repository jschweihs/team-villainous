<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/blog.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare Blog object
$blog = new Blog($db);
 
// Get id of entry to be edited
$data = json_decode(file_get_contents("php://input"));
 
// Set ID property of entry to be edited
$blog->id = $data->id;
 
// Set blog entry property values
$blog->title 	= $data->title;
$blog->user_id 	= $data->user_id;
$blog->preview 	= $data->preview;
$blog->content 	= $data->content;
$blog->updated 	= date('Y-m-d H:i:s');


if($blog->update()) {
	// Update the product
    http_response_code(200);
    echo json_encode(array("message" => "Blog entry was updated."));
} else {
	// Update failed
    http_response_code(503);
    echo json_encode(array("message" => "Unable to update blog."));
}
?>