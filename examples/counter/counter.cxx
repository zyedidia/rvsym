#include <backends/cxxrtl/cxxrtl.h>

#if defined(CXXRTL_INCLUDE_CAPI_IMPL) || \
    defined(CXXRTL_INCLUDE_VCD_CAPI_IMPL)
#include <backends/cxxrtl/cxxrtl_capi.cc>
#endif

#if defined(CXXRTL_INCLUDE_VCD_CAPI_IMPL)
#include <backends/cxxrtl/cxxrtl_vcd_capi.cc>
#endif

using namespace cxxrtl_yosys;

namespace cxxrtl_design {

// \top: 1
// \src: counter.sv:1.1-26.10
struct p_counter : public module {
	// \init: 0
	// \src: counter.sv:10.18-10.23
	wire<32> p_q__reg {0u};
	// \src: counter.sv:7.29-7.30
	/*output*/ value<32> p_q;
	// \src: counter.sv:6.26-6.29
	/*input*/ value<1> p_rst;
	// \src: counter.sv:6.21-6.24
	/*input*/ value<1> p_clk;
	value<1> prev_p_clk;
	bool posedge_p_clk() const {
		return !prev_p_clk.slice<0>().val() && p_clk.slice<0>().val();
	}
	p_counter() {}
	p_counter(adopt, p_counter other) {}

	void reset() override {
		*this = p_counter(adopt {}, std::move(*this));
	}

	bool eval() override;
	bool commit() override;
}; // struct p_counter

bool p_counter::eval() {
	bool converged = true;
	bool posedge_p_clk = this->posedge_p_clk();
	// cells $procdff$9 $procmux$7 $ternary$counter.sv:22$4 $eq$counter.sv:22$2 $add$counter.sv:22$3
	if (posedge_p_clk) {
		p_q__reg.next = (p_rst ? value<32>{0u} : (eq_uu<1>(p_q__reg.curr, value<32>{0x2710u}) ? value<32>{0u} : add_uu<32>(p_q__reg.curr, value<32>{0x1u})));
	}
	// connection
	p_q = p_q__reg.curr;
	return converged;
}

bool p_counter::commit() {
	bool changed = false;
	if (p_q__reg.commit()) changed = true;
	prev_p_clk = p_clk;
	return changed;
}

} // namespace cxxrtl_design

