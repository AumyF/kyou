{ pkgs ? import <nixpkgs> { } }: pkgs.buildGoModule rec {
  pname = "kyou";
  version = "0.1.0";

  src = ./.;

  vendorSha256 = "txbY5x88iLY3wtZmDuiVUo+s7DFa4U/8JvCNDfDcq1A=";
}
