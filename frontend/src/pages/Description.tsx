import React, { useEffect } from 'react';
import * as THREE from 'three';

const DescriptionPage: React.FC = () => {
  useEffect(() => {
    const scene = new THREE.Scene();
    const camera = new THREE.PerspectiveCamera(75, window.innerWidth / 400, 0.1, 1000);
    const renderer = new THREE.WebGLRenderer();
    renderer.setSize(window.innerWidth, 400);
    document.getElementById('three-container')?.appendChild(renderer.domElement);

    const geometry = new THREE.BoxGeometry();
    const material = new THREE.MeshBasicMaterial({ color: 0x00ffff });
    const cube = new THREE.Mesh(geometry, material);
    scene.add(cube);

    camera.position.z = 5;

    const animate = () => {
      requestAnimationFrame(animate);
      cube.rotation.x += 0.01;
      cube.rotation.y += 0.01;
      renderer.render(scene, camera);
    };

    animate();
  }, []);

  return (
    <div className="p-4">
      <h2 className="text-2xl font-semibold mb-2">Blockplain Description</h2>
      <p>This is a visual representation of how blocks interact in a 2D modular plane.</p>
      <div id="three-container" className="mt-4" />
    </div>
  );
};

export default DescriptionPage;
