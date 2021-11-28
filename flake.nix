{
  description = "A very basic flake";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system}; in
      rec {
        packages = flake-utils.lib.flattenTree
          {
            kyou = pkgs.callPackage ./default.nix { };
          };
        defaultPackage = packages.kyou;
        apps.kyou = flake-utils.lib.mkApp { drv = packages.kyou; };
        defaultApp = apps.kyou;
        devShell = pkgs.callPackage ./shell.nix { };
      }
    );
}
