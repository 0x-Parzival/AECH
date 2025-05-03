import React, { useState } from 'react';

const AddBlockPage: React.FC = () => {
  const [x, setX] = useState('');
  const [y, setY] = useState('');
  const [data, setData] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: Call backend to add block
    alert(`Block added at (${x}, ${y}) with data: ${data}`);
  };

  return (
    <div className="p-8">
      <h1 className="text-3xl font-bold mb-4">Add Block</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <input type="number" placeholder="X" value={x} onChange={(e) => setX(e.target.value)} className="border p-2 rounded w-24" required />
        <input type="number" placeholder="Y" value={y} onChange={(e) => setY(e.target.value)} className="border p-2 rounded w-24" required />
        <input type="text" placeholder="Data" value={data} onChange={(e) => setData(e.target.value)} className="border p-2 rounded w-64" required />
        <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">Add Block</button>
      </form>
    </div>
  );
};

export default AddBlockPage;
