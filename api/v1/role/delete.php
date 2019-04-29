<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
// Include database and object file
include_once '../config/database.php';
include_once '../objects/role.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare product object
$role = new Role($db);
// Get product id
$data = json_decode(file_get_contents("php://input"));
// Set product id to be deleted
$role->id = $data->id;
 
// Delete the product
if($role->delete()) {
    http_response_code(200);
    echo json_encode(array("message" => "Role was deleted."));
} else {
    http_response_code(503);
    echo json_encode(array("message" => "Unable to delete role."));
}
?>