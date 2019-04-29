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
    empty($data->email) ||
    empty($data->password)
) {
    http_response_code(400);
    echo json_encode(array("message" => "Unable to login. Data is incomplete."));
} else {
	// Set user property values
    $user->email         	= $data->email;
    $user->password         = $data->password;

    if($user->login()) {
    	// Create JWT
	    $token = array();
	    $token['id'] = $user->id;
	    $jwt = JWT::encode($token, 'secret_server_key');

	    $response = array();
	    $response['jwt'] = $jwt;
	    $response['user'] = array(
	    	"id"                => $user->id,
	        "username"          => $user->username,
	        "password"          => $user->passwrd,
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

    	http_response_code(200);
    	// Return user info along with JWT
        echo json_encode($response); 
    } else {
    	http_response_code(503);
        echo json_encode(array("message" => "Unable to login. Credentials are invalid."));
    }
}
?>