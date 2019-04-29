<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/role.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare product object
$role = new Role($db);
 
// Get id of product to be edited
$data = json_decode(file_get_contents("php://input"));
 
// Set ID property of product to be edited
$role->id = $data->id;
 
// Set product property values
$role->name = $data->name;
 
// Update the product
if($role->update()) {
    // set response code - 200 ok
    http_response_code(200);
    // tell the user
    echo json_encode(array("message" => "Role was updated."));
}
// If unable to update the product, tell the user
else{
    // Set response code - 503 service unavailable
    http_response_code(503);
    // Tell the user
    echo json_encode(array("message" => "Unable to update role."));
}
?>