<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Headers: access");
header("Access-Control-Allow-Methods: GET");
header("Access-Control-Allow-Credentials: true");
header('Content-Type: application/json');
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/user.php';
 
// Get database connection
$database = new Database();
$db = $database->getConnection();
 
// Prepare product object
$user = new User($db);
 
// Set ID property of record to read
$user->id = isset($_GET['id']) ? $_GET['id'] : die();
 
// Read the details of product to be edited
$user->get();
 
if($user->username != null) {
    // Create array
    $user_arr = array(
        "id"                => $user->id,
        "username"          => $user->username,
        "password"          => $user->password,
        "email"             => $user->email,
        "f_name"            => $user->f_name,
        "m_name"            => $user->m_name,
        "l_name"            => $user->l_name,
        "title"             => $user->title,
        "address"           => $user->address,
        "city"              => $user->city,
        "province"          => $user->province,
        "zip"               => $user->zip,
        "country"           => $user->country,
        "birth_date"        => $user->birth_date,
        "description"       => $user->description,
        "role"              => $user->role,
        "status"            => $user->status,
        "facebook_url"      => $user->facebook_url,
        "twitter_url"       => $user->twitter_url,
        "instagram_url"     => $user->instagram_url,
        "twitch_url"        => $user->twitch_url,
        "youtube_url"       => $user->youtube_url,
        "other_url"         => $user->other_url,
        "ps4_gamertag"      => $user->ps4_gamertag,
        "xbox_gamertag"     => $user->xbox_gamertag,
        "steam_gamertag"    => $user->steam_gamertag,
        "created"           => $user->created,
        "updated"           => $user->updated,
    );
 
    // Set response code - 200 OK
    http_response_code(200);
    // Make it json format
    echo json_encode($user_arr);
}
 
else{
    // Set response code - 404 Not found
    http_response_code(404);
    // Tell the user product does not exist
    echo json_encode(array("message" => "User does not exist."));
}
?>