<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Headers: access");
header("Access-Control-Allow-Methods: GET");
header("Access-Control-Allow-Credentials: true");
header('Content-Type: application/json');
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/blog.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare Blog object
$blog = new Blog($db);
 
// Set ID property of record to read
$blog->id = isset($_GET['id']) ? $_GET['id'] : die();
 
// Read the details of entry
$blog->get();
 
if($blog->title != null) {
    // Create array
    $blog_arr = array(
        "id"        => $blog->id,
        "title"     => $blog->title,
        "user_id"   => $blog->user_id,
        "preview"   => $blog->preview,
        "content"   => $blog->content,
        "created"   => $blog->created,
        "updated"   => $blog->updated,
    );
 
    // Set response code - 200 OK
    http_response_code(200);
    // Make it json format
    echo json_encode($blog_arr);
} else {
    // Set response code - 404 Not found
    http_response_code(404);
    // Tell the blog product does not exist
    echo json_encode(array("message" => "Blog entry does not exist."));
}
?>