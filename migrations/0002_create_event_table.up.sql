CREATE TABLE event (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    PetID INT NOT NULL,
    Date DATETIME NOT NULL,
    Type VARCHAR(255) NOT NULL,
    Remark TEXT,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (PetID) REFERENCES pet(ID) ON DELETE CASCADE
);

-- Create index on PetID for better performance
CREATE INDEX idx_event_petid ON event(PetID);
CREATE INDEX idx_event_date ON event(Date);