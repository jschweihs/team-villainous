<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
header("Access-Control-Allow-Methods: POST");
header("Access-Control-Max-Age: 3600");
header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");
 
include_once '../config/database.php';
include_once '../objects/user.php';
 
$database = new Database();
$db = $database->getConnection();
 
$user = new User($db);

// Get posted data
$data = json_decode(file_get_contents("php://input"));

// Make sure data is not empty
if (
    empty($data->username) ||
    empty($data->email) ||
    empty($data->password) ||
    empty($data->f_name) ||
    empty($data->role)
) {
    http_response_code(400);
    echo json_encode(array("message" => "Unable to create user. Data is incomplete."));
} else {

    // Set user property values
    $user->username         = $data->username;
    $user->password         = password_hash($data->password, PASSWORD_BCRYPT);
    $user->email            = $data->email;
    $user->f_name           = $data->f_name;
    $user->m_name           = $data->m_name;      
    $user->l_name           = $data->l_name;
    $user->title            = $data->title;
    $user->address          = $data->address;
    $user->city             = $data->city;
    $user->province         = $data->province;
    $user->zip              = $data->zip;
    $user->country          = $data->country;
    $user->birth_date       = $data->birth_date;
    $user->description      = $data->description;
    $user->role             = $data->role;
    $user->status           = $data->status;
    $user->facebook_url     = $data->facebook_url;
    $user->twitter_url      = $data->twitter_url;
    $user->instagram_url    = $data->instagram_url;
    $user->twitch_url       = $data->twitch_url;
    $user->youtube_url      = $data->youtube_url;
    $user->other_url        = $data->other_url;
    $user->ps4_gamertag     = $data->ps4_gamertag;
    $user->xbox_gamertag    = $data->xbox_gamertag;
    $user->steam_gamertag   = $data->steam_gamertag;
    $user->created          = date('Y-m-d H:i:s');
    $user->updated          = $user->created;
 
    $id = $user->create();
    // Create the user
    if($id) {
        // Set response code - 201 created
        http_response_code(201);
        // Tell the user
        echo json_encode(array("message" => "User was created.", "id" => $id));
    } else {
        http_response_code(503);
        echo json_encode(array("message" => "Unable to create user."));
    }
}
?>