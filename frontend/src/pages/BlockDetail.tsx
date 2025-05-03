import React from 'react';
import { useParams } from 'react-router-dom';

const BlockDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();

  // Later: fetch actual block data using this id
  const dummyBlock = {
    id,
    x: 1,
    y: 2,
    data: 'Some transaction data...',
    timestamp: '2025-05-03 18:45:00',
  };

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold mb-4">Block Details</h1>
      <div className="bg-gray-100 p-6 rounded-xl shadow-lg w-fit">
        <p className="text-lg"><strong>ID:</strong> {dummyBlock.id}</p>
        <p className="text-lg"><strong>Position:</strong> ({dummyBlock.x}, {dummyBlock.y})</p>
        <p className="text-lg"><strong>Timestamp:</strong> {dummyBlock.timestamp}</p>
        <p className="text-lg"><strong>Data:</strong> {dummyBlock.data}</p>
      </div>
    </div>
  );
};

export default BlockDetailPage;
