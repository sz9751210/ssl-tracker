// src/components/CertificateDetail.js

import React from "react";

const CertificateDetail = ({ certificate }) => {
  return (
    <div className="mt-4 p-4 border rounded shadow">
      <h2 className="text-lg font-semibold">Certificate Detail</h2>
      <p className="mt-2">Domain: {certificate.Domain}</p>
      <p>Expiration Date: {certificate.ExpirationDate}</p>
      <p>Status: {certificate.Status}</p>
    </div>
  );
};

export default CertificateDetail;
