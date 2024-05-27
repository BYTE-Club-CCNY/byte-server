{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  packages = with pkgs; [
    nodejs
    bun
    nodePackages.npm
    nodePackages.ts-node
  ];
}