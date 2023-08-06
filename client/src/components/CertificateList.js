// src/components/CertificateList.js

import React from "react";
import { ListGroup } from "react-bootstrap"; // 引入 Bootstrap 元素


const CertificateList = ({ certificates, onSelectCertificate }) => {
  return (
    <div className="mt-4 p-4 border rounded shadow">
      <h2 className="text-lg font-semibold">Certificates</h2>
      <ListGroup className="mt-2">
        {certificates.map((cert) => (
          <ListGroup.Item
            key={cert.Domain}
            onClick={() => onSelectCertificate(cert.Domain)}
            action
          >
            {cert.Domain} - {cert.Status}
          </ListGroup.Item>
        ))}
      </ListGroup>
    </div>
  );
};
export default CertificateList;
