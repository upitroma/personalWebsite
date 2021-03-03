// see https://www.skypack.dev/view/tsparticles for README

tsParticles.load("tsparticles", {
  fpsLimit: 60,
  background: {
    color: "#0f0f0f"
  },
  backgroundMode: {
    enable: true
  },
  particles: {
    color: {
      value: "#ffffff"
    },
    links: {
      color: "#6e6e6e",
      enable: true,
      triangles: {
        enable: false
      },
      distance: 120,
    },
    move: {
      enable: true,
      speed: 1
    },
    size: {
      value: 0,
      
    },
    collisions: {
        "enable": false
    },
    number: {
        density: {
          enable: true,
          area: 800,
          factor: 1200
        },
        limit: 0,
        value: 100
      },
  }
});
