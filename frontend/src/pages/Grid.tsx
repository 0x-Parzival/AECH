import React from 'react';
import { useNavigate } from 'react-router-dom';

type Block = {
  id: string;
  x: number;
  y: number;
};

const dummyBlocks: Block[] = [
  { id: 'A1', x: 0, y: 0 },
  { id: 'B2', x: 1, y: 0 },
  { id: 'C3', x: 2, y: 1 },
  { id: 'D4', x: 0, y: 2 },
];

const GridPage: React.FC = () => {
  const navigate = useNavigate();

  const handleClick = (block: Block) => {
    navigate(`/block/${block.id}`);
  };

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold mb-6">Block Grid</h1>
      <div className="grid grid-cols-4 gap-4">
        {dummyBlocks.map((block) => (
          <div
            key={block.id}
            onClick={() => handleClick(block)}
            className="bg-blue-500 text-white rounded-xl p-6 cursor-pointer hover:bg-blue-600 shadow-md"
          >
            <p className="text-lg font-semibold">ID: {block.id}</p>
            <p>
              ({block.x}, {block.y})
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default GridPage;
