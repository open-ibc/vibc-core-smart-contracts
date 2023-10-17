{
  description = "A development environment for vibc-core-sc";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-23.05";
    foundry.url = "github:shazow/foundry.nix/monthly"; # Use monthly branch for permanent releases
  };

  outputs = { self , nixpkgs, foundry, ... }: let
    system = "aarch64-darwin";
  in {
    devShells."${system}".default = let
      pkgs = import nixpkgs {
        inherit system;
        overlays = [
          (self: super: rec {
            nodejs = super.nodejs-18_x;
            pnpm = super.nodePackages.pnpm;
            yarn = (super.yarn.override { inherit nodejs; });
          })
        ] ++ [ foundry.overlay ];
      };
    in pkgs.mkShell {
      # create an environment with nodejs-18_x, pnpm, and yarn
      packages = with pkgs; [
        node2nix
        nodejs
        pnpm
        yarn
        ansible
        python3Packages.redis
        python3Packages.jmespath
        ansi2html
        foundry-bin
      ];

      shellHook = ''
        alias a='ansible-playbook'
        echo "node `${pkgs.nodejs}/bin/node --version`"
      '';
    };
  };
}
