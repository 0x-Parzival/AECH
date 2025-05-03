import React from 'react';

const GridPage: React.FC = () => {
  return (
    <div className="p-4">
      <h2 className="text-2xl font-semibold">Block Grid</h2>
      <div className="grid grid-cols-5 gap-4 mt-4">
        {/* Placeholder blocks */}
        {Array.from({ length: 25 }).map((_, idx) => (
          <div
            key={idx}
            className="bg-blue-100 p-4 border border-blue-300 rounded-lg"
          >
            Block {idx + 1}
          </div>
        ))}
      </div>
    </div>
  );
};
export default GridPage;
