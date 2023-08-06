// src/components/AddCertificateForm.js

import React, { useState } from "react";
import { Form, Button } from "react-bootstrap"; // 引入 Bootstrap 元素

const AddCertificateForm = ({ onAddCertificate }) => {
  const [domain, setDomain] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    const response = await fetch("http://localhost:8080/add-certificate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ domain }),
    });

    if (response.ok) {
      onAddCertificate();
      setDomain("");
    }
  };

  return (
    <div className="mt-4 p-4 border rounded shadow">
      <h2 className="text-lg font-semibold">Add Certificate</h2>
      <Form onSubmit={handleSubmit} className="mt-2">
        <Form.Group>
          <Form.Control
            type="text"
            value={domain}
            onChange={(e) => setDomain(e.target.value)}
            placeholder="Enter domain"
          />
        </Form.Group>
        <Button type="submit" variant="primary" className="mt-2">
          Add
        </Button>
      </Form>
    </div>
  );
};


export default AddCertificateForm;
