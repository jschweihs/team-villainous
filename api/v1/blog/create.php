<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
include_once '../config/database.php';
include_once '../objects/blog.php';
 
$database = new Database();
$db = $database->getConnection();
 
$blog = new Blog($db);

// Get posted data
$data = json_decode(file_get_contents("php://input"));

// Make sure data is not empty
if (
    empty($data->title) ||
    empty($data->user_id) ||
    empty($data->preview) ||
    empty($data->content)
) {
    http_response_code(400);
    echo json_encode(array("message" => "Unable to create entry. Data is incomplete."));
} else {

    // Set blog property values
    $blog->title    = $data->title;
    $blog->user_id  = $data->user_id;
    $blog->preview  = $data->preview;
    $blog->content  = $data->content;
    $blog->created  = date('Y-m-d H:i:s');
    $blog->updated  = $blog->created;
 
    $id = $blog->create();

    // Create the blog
    if($id) {
        // Set response code - 201 created
        http_response_code(201);
        // Tell the blog
        echo json_encode(array("message" => "Entry was created.", "id" => $id));
    } else {
        http_response_code(503);
        echo json_encode(array("message" => "Unable to create entry."));
    }
}
?>