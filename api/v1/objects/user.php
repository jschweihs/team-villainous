<?php

require_once '../lib/JWT.php';

class User {
 
    // Database connection and table name
    private $conn;
    private $table_name = "users";
 
    // Object properties
    public $id;
    public $username;
    public $password;
    public $email;
    public $f_name;
    public $m_name;
    public $l_name;
    public $title;
    public $address;
    public $city;
    public $province;
    public $zip;
    public $country;
    public $birth_date;
    public $description;
    public $role;
    public $status;
    public $facebook_url;
    public $twitter_url;
    public $instagram_url;
    public $twitch_url;
    public $youtube_url;
    public $other_url;
    public $ps4_gamertag;
    public $xbox_gamertag;
    public $steam_gamertag;
    public $created;
    public $updated;

 
    // Constructor with $db as database connection
    public function __construct($db){
        $this->conn = $db;
    }

    // Create user
    function create() {
        // Query to insert
        $query = "INSERT INTO " . $this->table_name . "
                SET
                    username=:username, 
                    password=:password,
                    email=:email,
                    f_name=:f_name,
                    m_name=:m_name,
                    l_name=:l_name,
                    title=:title,
                    address=:address,
                    city=:city,
                    province=:province,
                    zip=:zip,
                    country=:country,
                    birth_date=:birth_date,
                    description=:description,
                    role=:role,
                    status=:status,
                    facebook_url=:facebook_url,
                    twitter_url=:twitter_url,
                    instagram_url=:instagram_url,
                    twitch_url=:twitch_url,
                    youtube_url=:youtube_url,
                    other_url=:other_url,
                    ps4_gamertag=:ps4_gamertag,
                    xbox_gamertag=:xbox_gamertag,
                    steam_gamertag=:steam_gamertag,
                    created=:created,
                    updated=:updated";
     
        // Prepare query
        $stmt = $this->conn->prepare($query);
     
        // Sanitize
        $this->username         = htmlspecialchars(strip_tags($this->username));
        $this->email            = htmlspecialchars(strip_tags($this->email));
        $this->f_name           = htmlspecialchars(strip_tags($this->f_name));
        $this->m_name           = htmlspecialchars(strip_tags($this->m_name));
        $this->l_name           = htmlspecialchars(strip_tags($this->l_name));
        $this->title            = htmlspecialchars(strip_tags($this->title));
        $this->address          = htmlspecialchars(strip_tags($this->address));
        $this->city             = htmlspecialchars(strip_tags($this->city));
        $this->province         = htmlspecialchars(strip_tags($this->province));
        $this->zip              = htmlspecialchars(strip_tags($this->zip));
        $this->country          = htmlspecialchars(strip_tags($this->country));
        $this->birth_date       = htmlspecialchars(strip_tags($this->birth_date));
        $this->description      = htmlspecialchars(strip_tags($this->description));
        $this->facebook_url     = htmlspecialchars(strip_tags($this->facebook_url));
        $this->twitter_url      = htmlspecialchars(strip_tags($this->twitter_url));
        $this->instagram_url    = htmlspecialchars(strip_tags($this->instagram_url));
        $this->twitch_url       = htmlspecialchars(strip_tags($this->twitch_url));
        $this->youtube_url      = htmlspecialchars(strip_tags($this->youtube_url));
        $this->other_url        = htmlspecialchars(strip_tags($this->other_url));
        $this->ps4_gamertag     = htmlspecialchars(strip_tags($this->ps4_gamertag));
        $this->xbox_gamertag    = htmlspecialchars(strip_tags($this->xbox_gamertag));
        $this->steam_gamertag   = htmlspecialchars(strip_tags($this->steam_gamertag));
        $this->created          = htmlspecialchars(strip_tags($this->created));
        $this->updated          = htmlspecialchars(strip_tags($this->updated));

        // Bind values
        $stmt->bindParam(":username", $this->username);
        $stmt->bindParam(":password", $this->password);
        $stmt->bindParam(":email", $this->email);
        $stmt->bindParam(":f_name", $this->f_name);
        $stmt->bindParam(":m_name", $this->m_name);
        $stmt->bindParam(":l_name", $this->l_name);
        $stmt->bindParam(":title", $this->title);
        $stmt->bindParam(":address", $this->address);
        $stmt->bindParam(":city", $this->city);
        $stmt->bindParam(":province", $this->province);
        $stmt->bindParam(":zip", $this->zip);
        $stmt->bindParam(":country", $this->country);
        $stmt->bindParam(":birth_date", $this->birth_date);
        $stmt->bindParam(":description", $this->description);
        $stmt->bindParam(":role", $this->role);
        $stmt->bindParam(":status", $this->status);
        $stmt->bindParam(":facebook_url", $this->facebook_url);
        $stmt->bindParam(":twitter_url", $this->twitter_url);
        $stmt->bindParam(":instagram_url", $this->instagram_url);
        $stmt->bindParam(":twitch_url", $this->twitch_url);
        $stmt->bindParam(":youtube_url", $this->youtube_url);
        $stmt->bindParam(":other_url", $this->other_url);
        $stmt->bindParam(":ps4_gamertag", $this->ps4_gamertag);
        $stmt->bindParam(":xbox_gamertag", $this->xbox_gamertag);
        $stmt->bindParam(":steam_gamertag", $this->steam_gamertag);
        $stmt->bindParam(":created", $this->created);
        $stmt->bindParam(":updated", $this->updated);

        // Execute query
        if($stmt->execute()) {
            return $this->conn->lastInsertId();
        }
        return false;   
    }

    // Read users
    function getAll() {
        // Select all query
        $query = "SELECT * FROM " . $this->table_name . " ORDER BY created DESC";
        // Prepare query statement
        $stmt = $this->conn->prepare($query);
        // Execute query
        $stmt->execute();
        return $stmt;
    }

    // Get user by id
    function get() {
        // Query to read single record
        $query = "SELECT * FROM " . $this->table_name . " WHERE id = ? LIMIT 0,1";
        // Prepare query statement
        $stmt = $this->conn->prepare( $query );
        // Bind id of user to be updated
        $stmt->bindParam(1, $this->id);
        // Execute query
        $stmt->execute();
        // Get retrieved row
        $row = $stmt->fetch(PDO::FETCH_ASSOC);
     
        // Set values to object properties
        $this->username         = $row['username'];
        $this->password         = $row['password'];
        $this->email            = $row['email'];
        $this->f_name           = $row['f_name'];
        $this->m_name           = $row['m_name'];
        $this->l_name           = $row['l_name'];
        $this->title            = $row['title'];
        $this->address          = $row['address'];
        $this->city             = $row['city'];
        $this->province         = $row['province'];
        $this->zip              = $row['zip'];
        $this->country          = $row['country'];
        $this->birth_date       = $row['birth_date'];
        $this->description      = $row['description'];
        $this->role             = $row['role'];
        $this->status           = $row['status'];
        $this->facebook_url     = $row['facebook_url'];
        $this->twitter_url      = $row['twitter_url'];
        $this->instagram_url    = $row['instagram_url'];
        $this->twitch_url       = $row['twitch_url'];
        $this->youtube_url      = $row['youtube_url'];
        $this->other_url        = $row['other_url'];
        $this->ps4_gamertag     = $row['ps4_gamertag'];
        $this->xbox_gamertag    = $row['xbox_gamertag'];
        $this->steam_gamertag   = $row['steam_gamertag'];
        $this->created          = $row['created'];
        $this->updated          = $row['updated'];

    }

    // Update the user
    function update() {
        // Update query
        $query = "UPDATE " . $this->table_name . " SET
                    username=:username,
                    email=:email,
                    f_name=:f_name,
                    m_name=:m_name,
                    l_name=:l_name,
                    title=:title,
                    address=:address,
                    city=:city,
                    province=:province,
                    zip=:zip,
                    country=:country,
                    birth_date=:birth_date,
                    description=:description,
                    role=:role,
                    status=:status,
                    facebook_url=:facebook_url,
                    twitter_url=:twitter_url,
                    instagram_url=:instagram_url,
                    twitch_url=:twitch_url,
                    youtube_url=:youtube_url,
                    other_url=:other_url,
                    ps4_gamertag=:ps4_gamertag,
                    xbox_gamertag=:xbox_gamertag,
                    steam_gamertag=:steam_gamertag,
                    updated=:updated
                WHERE
                    id=:id";
                    
        // Prepare query statement
        $stmt = $this->conn->prepare($query);
     
        // Sanitize
        $this->username         = htmlspecialchars(strip_tags($this->username));
        $this->email            = htmlspecialchars(strip_tags($this->email));
        $this->f_name           = htmlspecialchars(strip_tags($this->f_name));
        $this->m_name           = htmlspecialchars(strip_tags($this->m_name));
        $this->l_name           = htmlspecialchars(strip_tags($this->l_name));
        $this->title            = htmlspecialchars(strip_tags($this->title));
        $this->address          = htmlspecialchars(strip_tags($this->address));
        $this->city             = htmlspecialchars(strip_tags($this->city));
        $this->province         = htmlspecialchars(strip_tags($this->province));
        $this->zip              = htmlspecialchars(strip_tags($this->zip));
        $this->country          = htmlspecialchars(strip_tags($this->country));
        $this->birth_date       = htmlspecialchars(strip_tags($this->birth_date));
        $this->description      = htmlspecialchars(strip_tags($this->description));
        $this->facebook_url     = htmlspecialchars(strip_tags($this->facebook_url));
        $this->twitter_url      = htmlspecialchars(strip_tags($this->twitter_url));
        $this->instagram_url    = htmlspecialchars(strip_tags($this->instagram_url));
        $this->twitch_url       = htmlspecialchars(strip_tags($this->twitch_url));
        $this->youtube_url      = htmlspecialchars(strip_tags($this->youtube_url));
        $this->other_url        = htmlspecialchars(strip_tags($this->other_url));
        $this->ps4_gamertag     = htmlspecialchars(strip_tags($this->ps4_gamertag));
        $this->xbox_gamertag    = htmlspecialchars(strip_tags($this->xbox_gamertag));
        $this->steam_gamertag   = htmlspecialchars(strip_tags($this->steam_gamertag));
        $this->updated          = htmlspecialchars(strip_tags($this->updated));

        // Bind new values
        $stmt->bindParam(':id', $this->id);
        $stmt->bindParam(":username", $this->username);
        $stmt->bindParam(":email", $this->email);
        $stmt->bindParam(":f_name", $this->f_name);
        $stmt->bindParam(":m_name", $this->m_name);
        $stmt->bindParam(":l_name", $this->l_name);
        $stmt->bindParam(":title", $this->title);
        $stmt->bindParam(":address", $this->address);
        $stmt->bindParam(":city", $this->city);
        $stmt->bindParam(":province", $this->province);
        $stmt->bindParam(":zip", $this->zip);
        $stmt->bindParam(":country", $this->country);
        $stmt->bindParam(":birth_date", $this->birth_date);
        $stmt->bindParam(":description", $this->description);
        $stmt->bindParam(":role", $this->role);
        $stmt->bindParam(":status", $this->status);
        $stmt->bindParam(":facebook_url", $this->facebook_url);
        $stmt->bindParam(":twitter_url", $this->twitter_url);
        $stmt->bindParam(":instagram_url", $this->instagram_url);
        $stmt->bindParam(":twitch_url", $this->twitch_url);
        $stmt->bindParam(":youtube_url", $this->youtube_url);
        $stmt->bindParam(":other_url", $this->other_url);
        $stmt->bindParam(":ps4_gamertag", $this->ps4_gamertag);
        $stmt->bindParam(":xbox_gamertag", $this->xbox_gamertag);
        $stmt->bindParam(":steam_gamertag", $this->steam_gamertag);
        $stmt->bindParam(":updated", $this->updated);

        // Execute the query
        if($stmt->execute()) {
            return true;
        }
        return false;
    }

    // Delete the user
    function delete() {
        $query = "DELETE FROM " . $this->table_name . " WHERE id = ?";
        $stmt = $this->conn->prepare($query);
        $this->id=htmlspecialchars(strip_tags($this->id));
        $stmt->bindParam(1, $this->id);

        if($stmt->execute()){
            return true;
        }
        return false;
         
    }

    // Login
    function login() {
        // Get users information
        $query = "SELECT id, email, password, f_name, m_name, l_name, title, address, city, province, zip, country, birth_date, description, role, status, facebook_url, twitter_url, instagram_url, twitch_url, youtube_url, other_url, ps4_gamertag, xbox_gamertag, steam_gamertag
            FROM " . $this->table_name . "
            WHERE email = ?
            LIMIT 0,1";

        $stmt = $this->conn->prepare($query);

        $this->email=htmlspecialchars(strip_tags($this->email));
        if($this->password != htmlspecialchars(strip_tags($this->password))) {
            // Avoid injection
            return;
        }

        $stmt->bindParam(1, $this->email);
        $stmt->execute();

        if($stmt->rowCount() == 0) {
            // No user was found
            return;
        } else {
            $row = $stmt->fetch(PDO::FETCH_ASSOC);
            if(!password_verify($this->password, $row['password'])) {
                // Password does not match
                return;
            } else {
                // User found with correct password
                $this->id               = $row['id'];
                $this->email            = $row['email'];
                $this->f_name           = $row['f_name'];
                $this->m_name           = $row['m_name'];
                $this->l_name           = $row['l_name'];
                $this->title            = $row['title'];
                $this->address          = $row['address'];
                $this->city             = $row['city'];
                $this->province         = $row['province'];
                $this->zip              = $row['zip'];
                $this->country          = $row['country'];
                $this->birth_date       = $row['birth_date'];
                $this->description      = $row['description'];
                $this->role             = $row['role'];
                $this->status           = $row['status'];
                $this->facebook_url     = $row['facebook_url'];
                $this->twitter_url      = $row['twitter_url'];
                $this->instagram_url    = $row['instagram_url'];
                $this->twitch_url       = $row['twitch_url'];
                $this->youtube_url      = $row['youtube_url'];
                $this->other_url        = $row['ther_url'];
                $this->ps4_gamertag     = $row['ps4_gamertag'];
                $this->xbox_gamertag    = $row['xbox_gamertag'];
                $this->steam_gamertag   = $row['steam_gamertag'];

                return true;
            }  
        }
    }
}