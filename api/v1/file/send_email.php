<?php

// header("Access-Control-Allow-Origin: *");
// header("Content-Type: application/json; charset=UTF-8");
// header("Access-Control-Allow-Methods: POST");
// header("Access-Control-Max-Age: 3600");
// header("Access-Control-Allow-Headers: Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With");

// if($_POST["submit"]) {

	$data = json_decode(file_get_contents("php://input"));

    $recipient		= "villainousteam2018@gmail.com";
    $subject		= $data->category;
    $sender 		= $data->name;
    $sender_email	= $data->address;
    $message		= $data->message;

    $mail_body 		= "Name: $sender\nEmail: $sender_email\n\n$message";

	// $header = "MIME-Version: 1.0" . "\r\n";
	// $header .= "Content-type:text/html;charset=UTF-8" . "\r\n";
// 	$headers = "From: webmaster@example.com" . "\r\n" .
// "CC: somebodyelse@example.com";
    // $mail_body = "test";
    // $header = "test";

    print_r($header);

    if(mail($recipient, $subject, $mail_body)) {
    	print_r("success");
    } else {
    	print_r("failure");
    }
// } else {
// 	print_r("No submission");
// }

?>