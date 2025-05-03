import React from 'react';
import { Link } from 'react-router-dom';

const HomePage: React.FC = () => {
  return (
    <div className="p-8 text-center">
      <h1 className="text-4xl font-bold mb-4">Welcome to Blockplain</h1>
      <p className="mb-6">Explore the modular 2D blockchain system visually</p>
      <div className="flex flex-wrap justify-center gap-4">
        <Link to="/grid" className="bg-blue-500 text-white px-4 py-2 rounded-xl">View Grid</Link>
        <Link to="/add-block" className="bg-green-500 text-white px-4 py-2 rounded-xl">Add Block</Link>
        <Link to="/simulate" className="bg-purple-500 text-white px-4 py-2 rounded-xl">Simulate Tx</Link>
        <Link to="/description" className="bg-gray-700 text-white px-4 py-2 rounded-xl">2D Description</Link>
      </div>
    </div>
  );
};

export default HomePage;
