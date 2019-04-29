<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Headers: access");
header("Access-Control-Allow-Methods: GET");
header("Access-Control-Allow-Credentials: true");
header('Content-Type: application/json');
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/role.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare product object
$role = new Role($db);
 
// Set ID property of record to read
$role->id = isset($_GET['id']) ? $_GET['id'] : die();
 
// Read the details of product to be edited
$role->get();
 
if($role->name != null) {
    // Create array
    $role_arr = array(
        "id"        => $role->id,
        "name"      => $role->name,
    );
 
    // Set response code - 200 OK
    http_response_code(200);
    // Make it json format
    echo json_encode($role_arr);
}
 
else{
    // Set response code - 404 Not found
    http_response_code(404);
    // Tell the user product does not exist
    echo json_encode(array("message" => "Role does not exist."));
}
?>