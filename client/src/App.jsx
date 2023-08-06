// src/App.js

import React, { useState, useEffect } from "react";
import AddCertificateForm from "./components/AddCertificateForm";
import CertificateList from "./components/CertificateList";
import CertificateDetail from "./components/CertificateDetail";

const App = () => {
  const [certificates, setCertificates] = useState([]);
  const [selectedCertificate, setSelectedCertificate] = useState(null);

  const fetchCertificates = async () => {
    const response = await fetch("http://localhost:8080/certificates");
    if (response.ok) {
      const data = await response.json();
      setCertificates(data);
    }
  };

  useEffect(() => {
    fetchCertificates();
  }, []);

  const handleAddCertificate = () => {
    fetchCertificates();
  };

  const handleSelectCertificate = (domain) => {
    const selected = certificates.find((cert) => cert.Domain === domain);
    setSelectedCertificate(selected);
  };

  return (
    <div>
      <h1>SSL Certificate Monitoring</h1>
      <AddCertificateForm onAddCertificate={handleAddCertificate} />
      <CertificateList
        certificates={certificates}
        onSelectCertificate={handleSelectCertificate}
      />
      {selectedCertificate && (
        <CertificateDetail certificate={selectedCertificate} />
      )}
    </div>
  );
};

export default App;
