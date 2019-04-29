<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
include_once '../config/database.php';
include_once '../objects/role.php';
 
$database = new Database();
$db = $database->getConnection();
$role = new Role($db);
 
// Get posted data
$data = json_decode(file_get_contents("php://input"));

// Make sure data is not empty
if (
    empty($data->name)
) {
    http_response_code(400);
    echo json_encode(array("message" => "Unable to create role. Data is incomplete."));
} else {
    // Set user property values
    $role->name = $data->name;
 
    $id = $role->create();
    // Create the user
    if($id) {
        // Set response code - 201 created
        http_response_code(201);
        // Tell the user
        echo json_encode(array("message" => "Role was created.", "id" => $id));
    } else {
        http_response_code(503);
        echo json_encode(array("message" => "Unable to create role."));
    }
}
?>