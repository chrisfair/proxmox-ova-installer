# proxmox-ova-installer

**proxmox-ova-installer** is a Go-based CLI tool that runs directly on a Proxmox VE host to import and export OVA virtual appliance packages.

This tool simplifies the deployment and packaging of virtual machines by automating OVA unpacking, disk conversion, and VM creation from within the Proxmox environment.

---

## ✨ Features

- ✅ Run directly on Proxmox — no external control node required  
- ✅ Extract `.ova` packages and deploy VMs using native Proxmox tools  
- ✅ Automatically convert `.vmdk` to `qcow2` using `qemu-img`  
- ✅ Generate `.ovf` and `.mf` files from existing VMs  
- ✅ Package a VM into a single `.ova` file for export or distribution  
- ✅ CLI-first design with clear flags and scripting support  

---

## 🧱 Project Structure

```
cmd/                # CLI binaries: import/export entrypoints  
internal/           # Core logic: unpacking, conversion, Proxmox interfacing  
pkg/                # Optional reusable libraries  
test/               # Integration/system tests  
Makefile            # Build and test automation  
LICENSE  
README.md  
go.mod  
```

---

## 🔧 Installation

Build from source directly on your Proxmox server:

```bash
git clone https://github.com/yourname/proxmox-ova-installer.git
cd proxmox-ova-installer
make build
```

Place the resulting binaries in your `PATH`, or run them from `./bin`.

---

## 🚀 Usage

### Import an OVA into Proxmox

```bash
sudo ./bin/ovaimport \
  --ova /root/images/myvm.ova \
  --vmid 105 \
  --storage local-lvm \
  --node $(hostname)
```

### Export an existing VM to OVA

```bash
sudo ./bin/ovaexport \
  --vmid 105 \
  --storage local-lvm \
  --out /root/export/myvm.ova \
  --node $(hostname)
```

---

## 📦 Requirements

- Run as `root` or a user with `qm`, `qemu-img`, and storage access  
- Proxmox VE 7.0+ with access to `local` or `local-lvm` storage  
- `qemu-img` installed (already present in Proxmox by default)  

---

## 🧪 Testing

Run all unit tests:

```bash
make test
```

Run integration tests (may require a test VM ID):

```bash
go test ./test/integration/...
```

---

## 📋 Roadmap

- [ ] Automatic VM config inference from `.ovf`  
- [ ] Streamable disk handling for large OVA files  
- [ ] Export Proxmox template VMs  
- [ ] Add `--dry-run` mode  
- [ ] Add `--force` overwrite option  

---

## 📝 License

MIT License — see [LICENSE](./LICENSE) for details.

---

## 🙏 Acknowledgments

- [Proxmox VE Documentation](https://pve.proxmox.com/)  
- [`qemu-img`](https://manpages.debian.org/qemu-utils/qemu-img.1.en.html)  
- [OVF Specification](https://www.dmtf.org/standards/ovf)

