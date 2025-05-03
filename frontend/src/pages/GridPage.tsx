import React from 'react';

const GridPage: React.FC = () => {
  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Blockplain Grid</h1>
      <div className="grid grid-cols-5 gap-4">
        {[...Array(25)].map((_, idx) => (
          <div key={idx} className="bg-gray-200 p-4 rounded shadow">
            Block at [{Math.floor(idx / 5)}, {idx % 5}]
          </div>
        ))}
      </div>
    </div>
  );
};

export default GridPage;
