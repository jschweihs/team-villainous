<?php

class Blog {
 
    // Database connection and table name
    private $conn;
    private $table_name = "blog";
 
    // Object properties
    public $id;
    public $title;
    public $user_id;
    public $preview;
    public $content;
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
                    title=:title,
                    user_id=:user_id,
                    preview=:preview,
                    content=:content,
                    created=:created,
                    updated=:updated";
     
        // Prepare query
        $stmt = $this->conn->prepare($query);

        // Sanitize
        $this->title    = htmlspecialchars(strip_tags($this->title));
        $this->user_id  = htmlspecialchars(strip_tags($this->user_id));
        $this->preview  = htmlspecialchars(strip_tags($this->preview));
        $this->created  = htmlspecialchars(strip_tags($this->created));
        $this->updated  = htmlspecialchars(strip_tags($this->updated));

        // Bind values
        $stmt->bindParam(":title", $this->title);
        $stmt->bindParam(":user_id", $this->user_id);
        $stmt->bindParam(":preview", $this->preview);
        $stmt->bindParam(":content", $this->content);
        $stmt->bindParam(":created", $this->created);
        $stmt->bindParam(":updated", $this->updated);

        // Execute query
        if($stmt->execute()) {
            return $this->conn->lastInsertId();
        }
        return false;   
    }

    // Read blog entries
    function getAll() {
        // Select all query
        $query = "SELECT blog.*, users.username FROM " . $this->table_name . " LEFT JOIN users ON blog.user_id  = users.id ORDER BY created DESC";
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
        $this->title    = $row['title'];
        $this->user_id  = $row['user_id'];
        $this->preview  = $row['preview'];
        $this->content  = $row['content'];
        $this->created  = $row['created'];
        $this->updated  = $row['updated'];

    }

    // Update the user
    function update() {
        // Update query
        $query = "UPDATE " . $this->table_name . " SET
                    title=:title,
                    user_id=:user_id,
                    preview=:preview,
                    content=:content,
                    updated=:updated
                WHERE
                    id=:id";

        // Prepare query statement
        $stmt = $this->conn->prepare($query);
     
        // Sanitize
        $this->title    = htmlspecialchars(strip_tags($this->title));
        $this->user_id  = htmlspecialchars(strip_tags($this->user_id));
        $this->preview  = htmlspecialchars(strip_tags($this->preview));
        $this->updated  = htmlspecialchars(strip_tags($this->updated));

        // Bind new values
        $stmt->bindParam(':id', $this->id);
        $stmt->bindParam(":title", $this->title);
        $stmt->bindParam(":user_id", $this->user_id);
        $stmt->bindParam(":preview", $this->preview);
        $stmt->bindParam(":content", $this->content);
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
}