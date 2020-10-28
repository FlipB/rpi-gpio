with import <nixpkgs> {};


mkShell {
  buildInputs = [
    go_1_15
    vgo2nix
    (import ./default.nix { inherit pkgs; })
  ];
}
