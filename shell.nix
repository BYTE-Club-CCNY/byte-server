{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  packages = with pkgs; [
    nodejs
    bun
    nodePackages.nodemon
    nodePackages.npm
    nodePackages.ts-node
  ];
}