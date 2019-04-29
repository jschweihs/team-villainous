<?php

class Role {
 
    // Database connection and table name
    private $conn;
    private $table_name = "roles";
 
    // Object properties
    public $id;
    public $name;
 
    // Constructor with $db as database connection
    public function __construct($db){
        $this->conn = $db;
    }

    // Create user
    function create() {
        // Query to insert
        $query = "INSERT INTO " . $this->table_name . "
                SET name=:name";
     
        // Prepare query
        $stmt = $this->conn->prepare($query);

        // Sanitize
        $this->name = htmlspecialchars(strip_tags($this->name));

        // Bind values
        $stmt->bindParam(":name", $this->name);
        // $stmt->bindParam(":created", $this->created);
        // $stmt->bindParam(":updated", $this->updated);

        // Execute query
        if($stmt->execute()) {
            return $this->conn->lastInsertId();
        }
        return false;   
    }

    // Read users
    function getAll() {
        // Select all query
        $query = "SELECT * FROM " . $this->table_name . " ORDER BY id ASC";
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
        // Bind id of role to be updated
        $stmt->bindParam(1, $this->id);
        // Execute query
        $stmt->execute();
        // Get retrieved row
        $row = $stmt->fetch(PDO::FETCH_ASSOC);
     
        // Set values to object properties
        $this->name     = $row['name'];
    }

    // Update role
    function update() {
        // Update query
        $query = "UPDATE " . $this->table_name . " SET name=:name WHERE id=:id";
     
        // Prepare query statement
        $stmt = $this->conn->prepare($query);
     
        // Sanitize
        $this->name = htmlspecialchars(strip_tags($this->name));

        // Bind new values
        $stmt->bindParam(':id', $this->id);
        $stmt->bindParam(':name', $this->name);

        // Execute the query
        return $stmt->execute();
    }

    // Delete the role
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