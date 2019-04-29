<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/role.php';
 
// Instantiate database and product object
$database = new Database();
$db = $database->getConnection();
 
// Initialize object
$role = new Role($db);
 
// Query users
$stmt = $role->getAll();
$num = $stmt->rowCount();
 
// Check if more than 0 records found
if($num>0){
 
    // products array
    $roles_arr=array();
    $roles_arr["roles"]=array();
 
    // retrieve our table contents
    // fetch() is faster than fetchAll()
    // http://stackoverflow.com/questions/2770630/pdofetchall-vs-pdofetch-in-a-loop
    while ($row = $stmt->fetch(PDO::FETCH_ASSOC)){
        // extract row
        // this will make $row['name'] to
        // just $name only
        extract($row);
 
        $role_row = array(
            "id"        => $id,
            "name" 		=> $name,
        );
 
        array_push($roles_arr["roles"], $role_row);
    }
 
    // Set response code - 200 OK
    http_response_code(200);
    // Show userss data in json format
    echo json_encode($roles_arr);
} else{
 
    // Set response code - 404 Not found
    http_response_code(404);
    // Tell the user no products found
    echo json_encode(
        array("message" => "No roles found.")
    );
}