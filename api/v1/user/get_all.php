<?php
// Required headers
header("Access-Control-Allow-Origin: *");
header("Content-Type: application/json; charset=UTF-8");
 
// Include database and object files
include_once '../config/database.php';
include_once '../objects/user.php';
 
// Instantiate database and product object
$database = new Database();
$db = $database->getConnection();
 
// Initialize object
$user = new User($db);
 
// Query users
$stmt = $user->getAll();
$num = $stmt->rowCount();
 
// Check if more than 0 records found
if($num>0){
 
    // products array
    $users_arr=array();
    $users_arr["records"]=array();
 
    // retrieve our table contents
    // fetch() is faster than fetchAll()
    // http://stackoverflow.com/questions/2770630/pdofetchall-vs-pdofetch-in-a-loop
    while ($row = $stmt->fetch(PDO::FETCH_ASSOC)){
        // extract row
        // this will make $row['name'] to
        // just $name only
        extract($row);
 
        $user_row = array(
            "id"                => $id,
            "username"          => $username,
            "password"          => $password,
            "email"             => $email,
            "f_name"            => $f_name,
            "m_name"            => $m_name,
            "l_name"            => $l_name,
            "title"             => $title,
            "address"           => $address,
            "city"              => $city,
            "province"          => $province,
            "zip"               => $zip,
            "country"           => $country,
            "birth_date"        => $birth_date,
            "description"       => $description,
            "role"              => $role,
            "status"            => $status,
            "facebook_url"      => $facebook_url,
            "twitter_url"       => $twitter_url,
            "instagram_url"     => $instagram_url,
            "twitch_url"        => $twitch_url,
            "youtube_url"       => $youtube_url,
            "other_url"         => $other_url,
            "ps4_gamertag"      => $ps4_gamertag,
            "xbox_gamertag"     => $xbox_gamertag,
            "steam_gamertag"    => $steam_gamertag,
            "created"           => $created,
            "updated"           => $updated
        );
 
        array_push($users_arr["records"], $user_row);
    }
 
    // Set response code - 200 OK
    http_response_code(200);
 
    // Show userss data in json format
    echo json_encode($users_arr);
} else{
 
    // Set response code - 404 Not found
    http_response_code(404);
 
    // Tell the user no products found
    echo json_encode(
        array("message" => "No users found.")
    );
}