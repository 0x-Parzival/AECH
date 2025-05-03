import React from 'react';
import { useParams } from 'react-router-dom';

const BlockDetailPage: React.FC = () => {
  const { id } = useParams();

  return (
    <div>
      <h1>Block Detail Page</h1>
      <p>Details of block {id}</p>
      {/* Display block details based on ID */}
    </div>
  );
};

export default BlockDetailPage;
