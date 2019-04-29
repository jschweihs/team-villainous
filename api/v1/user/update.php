<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/user.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare product object
$user = new User($db);
 
// Get id of product to be edited
$data = json_decode(file_get_contents("php://input"));
 
// Set ID property of product to be edited
$user->id = $data->id;
 
// Set product property values
$user->username 		= $data->username;
$user->email 			= $data->email;
$user->f_name 			= $data->f_name;
$user->m_name 			= $data->m_name;
$user->l_name 			= $data->l_name;
$user->title 			= $data->title;
$user->address 			= $data->address;
$user->city 			= $data->city;
$user->province 		= $data->province;
$user->zip 				= $data->zip;
$user->country 			= $data->country;
$user->birth_date 		= $data->birth_date;
$user->description 		= $data->description;
$user->role 			= $data->role;
$user->status 			= $data->status;
$user->facebook_url 	= $data->facebook_url;
$user->twitter_url		= $data->twitter_url;
$user->instagram_url 	= $data->instagram_url;
$user->twitch_url 		= $data->twitch_url;
$user->youtube_url 		= $data->youtube_url;
$user->other_url 		= $data->other_url;
$user->ps4_gamertag 	= $data->ps4_gamertag;
$user->xbox_gamertag 	= $data->xbox_gamertag;
$user->steam_gamertag	= $data->steam_gamertag;
$user->updated 			= date('Y-m-d H:i:s');


if($user->update()) {
	// Update the product
    http_response_code(200);
    echo json_encode(array("message" => "User was updated."));
} else {
	// Update failed
    http_response_code(503);
    echo json_encode(array("message" => "Unable to update user."));
}
?>