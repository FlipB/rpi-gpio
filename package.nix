{ stdenv, buildGoModule, fetchFromGitHub }:

# Assert go version
#assert lib.versionAtLeast go.version "1.13";

buildGoModule rec {
  name = "rpi-gpio";
  pname = "rpi-gpio";
  # version = "v1.0.0";
  # rev = "v${version}";
  rev = "master";

  src = ./.;

  # The hash of the output of the intermediate fetcher derivation
  vendorSha256 = "0sjjj9z1dhilhpc8pq4154czrb79z9cm044jvn75kxcjv6v5l2m5";

  # runVend runs the vend command to generate the vendor directory. This is useful if your code depends on c code and go mod tidy does not include the needed sources to build. 
  runVend = false;

  # Child packages to build
  subPackages = ["./cmd/rpi-gpio"];

  preBuild = ''
    export CGO_ENABLED=0
  '';

  meta = with stdenv.lib; {
    description = "CLI tool for setting GPIO pin states on raspberry pis";
    homepage = "https://github.com/flipb/rpi-gpio";
    platforms = platforms.linux;
    license = licenses.mit;
    maintainers = with maintainers; [ flipb ];
  };
}

