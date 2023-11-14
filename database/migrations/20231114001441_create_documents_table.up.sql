CREATE TABLE documents (
    id INT NOT NULL AUTO_INCREMENT,
    document_id VARCHAR(40) NOT NULL,
    user_id INT NOT NULL,

    risk_name VARCHAR(50) NOT NULL,

    fraud_schema VARCHAR(1000) NOT NULL,
    fraud_motive VARCHAR(1000) NOT NULL,
    fraud_technique VARCHAR(1000) NOT NULL,

    risk_source VARCHAR(1000) NOT NULL,
    root_cause VARCHAR(1000) NOT NULL,
    bispro_control_procedure VARCHAR(1000) NOT NULL,
    qualitative_impact VARCHAR(1000) NOT NULL,

    likehood_justification VARCHAR(1000) NOT NULL,
    impact_justification VARCHAR(1000) NOT NULL,

    strategy_agreement VARCHAR(1000) NOT NULL,
    strategy_recomendation VARCHAR(1000) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE (document_id)
)