<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/blog.php';
 
// Instantiate database and product object
$database = new Database();
$db = $database->getConnection();
 
// Initialize object
$blog = new Blog($db);
 
// Query blogs
$stmt = $blog->getAll();
$num = $stmt->rowCount();
 
// Check if more than 0 records found
if($num>0){
 
    // Entries array
    $blogs_arr=array();
    $blogs_arr["entries"]=array();

    while ($row = $stmt->fetch(PDO::FETCH_ASSOC)){

        extract($row);
 
        $blog_row = array(
            "id"        => $id,
            "title"     => $title,
            "user_id"   => $user_id,
            "username"  => $username,
            "preview"   => $preview,
            "content"   => $content,
            "created"   => $created,
            "updated"   => $updated
        );
 
        array_push($blogs_arr["entries"], $blog_row);
    }
 
    // Set response code - 200 OK
    http_response_code(200);
 
    // Show blogs data in json format
    echo json_encode($blogs_arr);
} else{
 
    // Set response code - 404 Not found
    http_response_code(404);
 
    // Tell the user no entriess found
    echo json_encode(
        array("message" => "No blogs found.")
    );
}