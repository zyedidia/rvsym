#include <cxx.h>

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
// \src: ./soc_top.sv:1.1-190.10
struct p_soc__top : public module {
	// \hdlname: gpio_unit gpo_reg
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:23.18-23.25
	wire<32> p_gpio__unit_2e_gpo__reg;
	// \hdlname: gpio_unit gpio_fsel_reg
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:24.18-24.31
	wire<32> p_gpio__unit_2e_gpio__fsel__reg;
	// \hdlname: gpio_unit wr_valid
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:25.11-25.19
	wire<1> p_gpio__unit_2e_wr__valid;
	// \hdlname: gpo_unit gpo_reg
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:26.19-26.26
	wire<32> p_gpo__unit_2e_gpo__reg;
	// \hdlname: gpo_unit wr_valid_reg
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:27.11-27.23
	wire<1> p_gpo__unit_2e_wr__valid__reg;
	// \hdlname: timer_unit timer_reg
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:24.18-24.27
	wire<32> p_timer__unit_2e_timer__reg;
	// \hdlname: timer_unit tick_reg
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:25.39-25.47
	wire<6> p_timer__unit_2e_tick__reg;
	// \hdlname: ram_unit read_word
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:38.18-38.27
	wire<32> p_ram__unit_2e_read__word;
	// \hdlname: ram_unit rd_valid
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:42.11-42.19
	wire<1> p_ram__unit_2e_rd__valid;
	// \hdlname: ram_unit wr_valid
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:42.21-42.29
	wire<1> p_ram__unit_2e_wr__valid;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__ADDR;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__DATA;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__EN;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__ADDR;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__DATA;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__EN;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__ADDR;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__DATA;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__EN;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__ADDR;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__DATA;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
	wire<32> i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__EN;
	// \init: 0
	// \hdlname: cpu_unit state_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:70.17-70.26
	wire<3> p_cpu__unit_2e_state__reg {0u};
	// \hdlname: cpu_unit alu_result_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:163.18-163.32
	wire<32> p_cpu__unit_2e_alu__result__reg;
	// \hdlname: cpu_unit store_data_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:164.18-164.32
	wire<32> p_cpu__unit_2e_store__data__reg;
	// \init: 0
	// \hdlname: cpu_unit fetch_unit pc_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:26.18-26.24|./core/cpu.sv:102.11-115.6
	wire<32> p_cpu__unit_2e_fetch__unit_2e_pc__reg {0u};
	// \hdlname: cpu_unit fetch_unit instr
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:21.33-21.38|./core/cpu.sv:102.11-115.6
	wire<32> p_cpu__unit_2e_fetch__unit_2e_instr;
	// \hdlname: cpu_unit decode_unit reg_file_unit r0_data
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:13.39-13.46|./core/cpu.sv:148.12-159.6
	wire<32> p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r0__data;
	// \hdlname: cpu_unit decode_unit reg_file_unit r1_data
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:13.48-13.55|./core/cpu.sv:148.12-159.6
	wire<32> p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:0.0-0.0|./core/cpu.sv:148.12-159.6
	wire<5> i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__ADDR;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:0.0-0.0|./core/cpu.sv:148.12-159.6
	wire<32> i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__DATA;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:0.0-0.0|./core/cpu.sv:148.12-159.6
	wire<32> i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__EN;
	// \hdlname: cpu_unit memory_unit mem_rdata_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:72.18-72.31|./core/cpu.sv:195.12-217.6
	wire<32> p_cpu__unit_2e_memory__unit_2e_mem__rdata__reg;
	// \src: ./soc_top.sv:23.22-23.32
	/*output*/ value<1> p_rgb__led0__b;
	// \src: ./soc_top.sv:22.22-22.32
	/*output*/ value<1> p_rgb__led0__g;
	// \src: ./soc_top.sv:21.22-21.32
	/*output*/ value<1> p_rgb__led0__r;
	// \src: ./soc_top.sv:19.21-19.28
	/*inout*/ value<1> p_gpio__a3;
	// \src: ./soc_top.sv:18.21-18.28
	/*inout*/ value<1> p_gpio__a2;
	// \src: ./soc_top.sv:17.21-17.28
	/*inout*/ value<1> p_gpio__a1;
	// \src: ./soc_top.sv:16.21-16.28
	/*inout*/ value<1> p_gpio__a0;
	// \src: ./soc_top.sv:15.21-15.28
	/*inout*/ value<1> p_gpio__13;
	// \src: ./soc_top.sv:14.21-14.28
	/*inout*/ value<1> p_gpio__12;
	// \src: ./soc_top.sv:13.21-13.28
	/*inout*/ value<1> p_gpio__11;
	// \src: ./soc_top.sv:12.21-12.28
	/*inout*/ value<1> p_gpio__10;
	// \src: ./soc_top.sv:11.21-11.27
	/*inout*/ value<1> p_gpio__9;
	// \src: ./soc_top.sv:10.21-10.27
	/*inout*/ value<1> p_gpio__6;
	// \src: ./soc_top.sv:9.21-9.27
	/*inout*/ value<1> p_gpio__5;
	// \src: ./soc_top.sv:8.21-8.27
	/*inout*/ value<1> p_gpio__1;
	// \src: ./soc_top.sv:7.21-7.27
	/*inout*/ value<1> p_gpio__0;
	// \src: ./soc_top.sv:5.21-5.28
	/*input*/ value<1> p_usr__btn;
	// \src: ./soc_top.sv:3.26-3.29
	/*input*/ value<1> p_rst;
	// \src: ./soc_top.sv:3.21-3.24
	/*input*/ value<1> p_clk;
	value<1> prev_p_clk;
	bool posedge_p_clk() const {
		return !prev_p_clk.slice<0>().val() && p_clk.slice<0>().val();
	}
	// \hdlname: ram_unit mem
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:32.18-32.21
	memory<32, 2048> memory_p_ram__unit_2e_mem {
		// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:0.0-0.0
		memory<32, 2048>::init<2> { 0, {
			value<32>{0x00000000u}, value<32>{0x00000042u},
		}},
	};
	// \hdlname: cpu_unit decode_unit reg_file_unit regs
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:16.28-16.32|./core/cpu.sv:148.12-159.6
	memory<32, 32> memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs {
		memory<32, 32>::init<32> { 0, {
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
			value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u}, value<32>{0x00000000u},
		}},
	};

	p_soc__top() {}
	p_soc__top(adopt, p_soc__top other) {}

	void reset() override {
		*this = p_soc__top(adopt {}, std::move(*this));
	}

	bool eval() override;
	bool commit() override;
}; // struct p_soc__top

bool p_soc__top::eval() {
	bool converged = true;
	bool posedge_p_clk = this->posedge_p_clk();
	value<5> i_procmux_24_1009__Y;
	value<32> i_procmux_24_1006__Y;
	value<32> i_procmux_24_1003__Y;
	value<1> i_procmux_24_1000__Y;
	value<2> i_procmux_24_796__CMP;
	value<2> i_procmux_24_795__CMP;
	value<3> i_procmux_24_786__CMP;
	value<3> i_procmux_24_785__CMP;
	value<1> i_procmux_24_745__Y;
	value<32> i_procmux_24_739__Y;
	value<1> i_procmux_24_727__Y;
	value<32> i_procmux_24_724__Y;
	value<3> i_procmux_24_608__Y;
	value<32> i_procmux_24_603__Y;
	value<32> i_procmux_24_600__Y;
	value<32> i_procmux_24_597__Y;
	value<32> i_procmux_24_594__Y;
	value<32> i_procmux_24_591__Y;
	value<32> i_procmux_24_588__Y;
	value<32> i_procmux_24_585__Y;
	value<32> i_procmux_24_582__Y;
	value<32> i_procmux_24_579__Y;
	value<32> i_procmux_24_576__Y;
	value<32> i_procmux_24_573__Y;
	value<32> i_procmux_24_570__Y;
	// \hdlname: gpio_unit clk
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:3.21-3.24
	value<1> p_gpio__unit_2e_clk;
	// \hdlname: gpio_unit rst
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:3.26-3.29
	value<1> p_gpio__unit_2e_rst;
	// \hdlname: gpio_unit gpio_inout
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:5.28-5.38
	value<32> p_gpio__unit_2e_gpio__inout;
	// \hdlname: gpio_unit i_rd
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:7.28-7.32
	value<1> p_gpio__unit_2e_i__rd;
	// \hdlname: gpio_unit i_addr
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:8.28-8.34
	value<32> p_gpio__unit_2e_i__addr;
	// \hdlname: gpio_unit i_wr
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:9.28-9.32
	value<1> p_gpio__unit_2e_i__wr;
	// \hdlname: gpio_unit i_data
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:11.28-11.34
	value<32> p_gpio__unit_2e_i__data;
	// \hdlname: gpio_unit gpio_addr
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:20.11-20.20
	value<1> p_gpio__unit_2e_gpio__addr;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:47.13-47.32
	value<1> i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:0.0-0.0
	value<1> i_flatten_5c_gpio__unit_2e__24_1_5c_rd__valid_5b_0_3a_0_5d_;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:0.0-0.0
	value<32> i_flatten_5c_gpio__unit_2e__24_1_5c_rd__data_5b_31_3a_0_5d_;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:55.13-55.39
	value<1> i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:57.22-57.50
	value<1> i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:59.22-59.50
	value<1> i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_59_24_33__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_37__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_38__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_39__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_40__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_41__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_42__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_43__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_44__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_45__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_46__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_47__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_48__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_49__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_50__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_51__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_52__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_53__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_54__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_55__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_56__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_57__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_58__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_59__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_60__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_61__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_62__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_63__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_64__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_65__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_66__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_67__Y;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	value<1> i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_68__Y;
	// \hdlname: gpo_unit clk
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:6.21-6.24
	value<1> p_gpo__unit_2e_clk;
	// \hdlname: gpo_unit rst
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:6.26-6.29
	value<1> p_gpo__unit_2e_rst;
	// \hdlname: gpo_unit i_addr
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:11.28-11.34
	value<32> p_gpo__unit_2e_i__addr;
	// \hdlname: gpo_unit i_wr
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:12.28-12.32
	value<1> p_gpo__unit_2e_i__wr;
	// \hdlname: gpo_unit i_data
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:14.28-14.34
	value<32> p_gpo__unit_2e_i__data;
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:41.13-41.38
	value<1> i_flatten_5c_gpo__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpo_2e_sv_3a_41_24_83__Y;
	// \hdlname: gpi_unit i_addr
	// \src: ./soc_top.sv:113.18-127.6|./sys/gpio/gpi.sv:11.28-11.34
	value<32> p_gpi__unit_2e_i__addr;
	// \hdlname: gpi_unit gpi_addr
	// \src: ./soc_top.sv:113.18-127.6|./sys/gpio/gpi.sv:23.11-23.19
	value<1> p_gpi__unit_2e_gpi__addr;
	// \src: ./soc_top.sv:113.18-127.6|./sys/gpio/gpi.sv:32.13-32.40
	value<1> i_flatten_5c_gpi__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpi_2e_sv_3a_32_24_379__Y;
	// \hdlname: timer_unit clk
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:6.21-6.24
	value<1> p_timer__unit_2e_clk;
	// \hdlname: timer_unit rst
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:6.26-6.29
	value<1> p_timer__unit_2e_rst;
	// \hdlname: timer_unit timer_addr
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:21.11-21.21
	value<1> p_timer__unit_2e_timer__addr;
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:43.13-43.31
	value<1> i_flatten_5c_timer__unit_2e__24_logic__and_24__2e__2f_sys_2f_timer_2e_sv_3a_43_24_93__Y;
	// \hdlname: ram_unit clk
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:13.21-13.24
	value<1> p_ram__unit_2e_clk;
	// \hdlname: ram_unit i_addr
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:16.28-16.34
	value<32> p_ram__unit_2e_i__addr;
	// \hdlname: ram_unit ram_addr
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:27.11-27.19
	value<1> p_ram__unit_2e_ram__addr;
	// \hdlname: ram_unit write_word
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:39.18-39.28
	value<32> p_ram__unit_2e_write__word;
	// \hdlname: ram_unit write_mask
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:40.17-40.27
	value<4> p_ram__unit_2e_write__mask;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:63.13-63.29
	value<1> i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:76.26-76.29
	value<32> i_flatten_5c_ram__unit_2e__24_memrd_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_76_24_163__DATA;
	// \hdlname: cpu_unit clk
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:3.21-3.24
	value<1> p_cpu__unit_2e_clk;
	// \hdlname: cpu_unit rst
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:3.26-3.29
	value<1> p_cpu__unit_2e_rst;
	// \hdlname: cpu_unit pc_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:99.18-99.24
	value<32> p_cpu__unit_2e_pc__reg;
	// \hdlname: cpu_unit imm
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:122.18-122.21
	value<32> p_cpu__unit_2e_imm;
	// \hdlname: cpu_unit mem_write
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:126.22-126.31
	value<1> p_cpu__unit_2e_mem__write;
	// \hdlname: cpu_unit next_pc
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:127.17-127.24
	value<2> p_cpu__unit_2e_next__pc;
	// \hdlname: cpu_unit wb_src
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:127.26-127.32
	value<2> p_cpu__unit_2e_wb__src;
	// \hdlname: cpu_unit ext_ctrl
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:128.17-128.25
	value<3> p_cpu__unit_2e_ext__ctrl;
	// \hdlname: cpu_unit pc_plus4
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:221.18-221.26
	value<32> p_cpu__unit_2e_pc__plus4;
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:222.23-222.37
	value<32> i_flatten_5c_cpu__unit_2e__24_add_24__2e__2f_core_2f_cpu_2e_sv_3a_222_24_327__Y;
	// \hdlname: cpu_unit bus_unit o_bus_data
	// \src: ./soc_top.sv:55.9-67.6|./core/bus.sv:27.29-27.39|./core/cpu.sv:38.9-64.6
	value<32> p_cpu__unit_2e_bus__unit_2e_o__bus__data;
	// \hdlname: cpu_unit bus_unit mem_write_req_valid
	// \src: ./soc_top.sv:55.9-67.6|./core/bus.sv:14.21-14.40|./core/cpu.sv:38.9-64.6
	value<1> p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid;
	// \hdlname: cpu_unit bus_unit mem_read_req_valid
	// \src: ./soc_top.sv:55.9-67.6|./core/bus.sv:10.21-10.39|./core/cpu.sv:38.9-64.6
	value<1> p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid;
	// \hdlname: cpu_unit bus_unit mem_read_res_data
	// \src: ./soc_top.sv:55.9-67.6|./core/bus.sv:9.29-9.46|./core/cpu.sv:38.9-64.6
	value<32> p_cpu__unit_2e_bus__unit_2e_mem__read__res__data;
	// \hdlname: cpu_unit bus_unit instr_read_req_valid
	// \src: ./soc_top.sv:55.9-67.6|./core/bus.sv:5.21-5.41|./core/cpu.sv:38.9-64.6
	value<1> p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid;
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:45.25-45.84|./core/cpu.sv:102.11-115.6
	value<3> i_flatten_5c_cpu__unit_2e__5c_fetch__unit_2e__24_ternary_24__2e__2f_core_2f_fetch_2e_sv_3a_45_24_316__Y;
	// \hdlname: cpu_unit fetch_unit state_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:9.27-9.36|./core/cpu.sv:102.11-115.6
	value<3> p_cpu__unit_2e_fetch__unit_2e_state__reg;
	// \hdlname: cpu_unit fetch_unit rst
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:7.26-7.29|./core/cpu.sv:102.11-115.6
	value<1> p_cpu__unit_2e_fetch__unit_2e_rst;
	// \hdlname: cpu_unit fetch_unit clk
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:7.21-7.24|./core/cpu.sv:102.11-115.6
	value<1> p_cpu__unit_2e_fetch__unit_2e_clk;
	// \hdlname: cpu_unit control_unit extract_imm_unit ctrl
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:72.17-74.6|./core/extract_imm.sv:6.27-6.31|./core/cpu.sv:131.13-142.6
	value<3> p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl;
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:97.18-97.37|./core/cpu.sv:131.13-142.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_97_24_304__Y;
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:96.18-96.47|./core/cpu.sv:131.13-142.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y;
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:95.13-95.35|./core/cpu.sv:131.13-142.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y;
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:65.18-65.49|./core/cpu.sv:131.13-142.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_65_24_286__Y;
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:62.13-62.31|./core/cpu.sv:131.13-142.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y;
	// \hdlname: cpu_unit control_unit is_rarith
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:27.22-27.31|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__rarith;
	// \hdlname: cpu_unit control_unit is_iarith
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:27.11-27.20|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__iarith;
	// \hdlname: cpu_unit control_unit is_store
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.66-26.74|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__store;
	// \hdlname: cpu_unit control_unit is_load
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.57-26.64|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__load;
	// \hdlname: cpu_unit control_unit is_branch
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.46-26.55|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__branch;
	// \hdlname: cpu_unit control_unit is_jalr
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.37-26.44|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__jalr;
	// \hdlname: cpu_unit control_unit is_jal
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.29-26.35|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__jal;
	// \hdlname: cpu_unit control_unit is_auipc
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.19-26.27|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__auipc;
	// \hdlname: cpu_unit control_unit is_lui
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:26.11-26.17|./core/cpu.sv:131.13-142.6
	value<1> p_cpu__unit_2e_control__unit_2e_is__lui;
	// \hdlname: cpu_unit control_unit funct3
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:19.17-19.23|./core/cpu.sv:131.13-142.6
	value<3> p_cpu__unit_2e_control__unit_2e_funct3;
	// \hdlname: cpu_unit control_unit op
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:18.17-18.19|./core/cpu.sv:131.13-142.6
	value<7> p_cpu__unit_2e_control__unit_2e_op;
	// \hdlname: cpu_unit control_unit instr
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:4.28-4.33|./core/cpu.sv:131.13-142.6
	value<32> p_cpu__unit_2e_control__unit_2e_instr;
	// \hdlname: cpu_unit decode_unit reg_file_unit clk
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:9.21-9.24|./core/cpu.sv:148.12-159.6
	value<1> p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_clk;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:25.13-25.33|./core/cpu.sv:148.12-159.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_logic__and_24__2e__2f_core_2f_reg__file_2e_sv_3a_25_24_210__Y;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:28.20-28.24|./core/cpu.sv:148.12-159.6
	value<32> i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_28_24_214__DATA;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:29.20-29.24|./core/cpu.sv:148.12-159.6
	value<32> i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_29_24_215__DATA;
	// \hdlname: cpu_unit execute_unit reg_extend_unit ctrl
	// \src: ./soc_top.sv:55.9-67.6|./core/execute.sv:42.12-46.6|./core/extend.sv:6.27-6.31|./core/cpu.sv:166.13-180.6
	value<3> p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl;
	// \hdlname: cpu_unit execute_unit alu_unit a
	// \src: ./soc_top.sv:55.9-67.6|./core/execute.sv:32.9-37.6|./core/alu.sv:3.28-3.29|./core/cpu.sv:166.13-180.6
	value<32> p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a;
	// \hdlname: cpu_unit execute_unit alu_unit b
	// \src: ./soc_top.sv:55.9-67.6|./core/execute.sv:32.9-37.6|./core/alu.sv:3.31-3.32|./core/cpu.sv:166.13-180.6
	value<32> p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b;
	// \hdlname: cpu_unit execute_unit alu_unit op
	// \src: ./soc_top.sv:55.9-67.6|./core/execute.sv:32.9-37.6|./core/alu.sv:4.27-4.29|./core/cpu.sv:166.13-180.6
	value<3> p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op;
	// \hdlname: cpu_unit execute_unit rd2
	// \src: ./soc_top.sv:55.9-67.6|./core/execute.sv:11.33-11.36|./core/cpu.sv:166.13-180.6
	value<32> p_cpu__unit_2e_execute__unit_2e_rd2;
	// \hdlname: cpu_unit memory_unit mem_read_extend_unit ctrl
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:76.12-80.6|./core/extend.sv:6.27-6.31|./core/cpu.sv:195.12-217.6
	value<3> p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl;
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:61.13-61.55|./core/cpu.sv:195.12-217.6
	value<1> i_flatten_5c_cpu__unit_2e__5c_memory__unit_2e__24_logic__and_24__2e__2f_core_2f_memory_2e_sv_3a_61_24_359__Y;
	// \hdlname: cpu_unit memory_unit write_req_outstanding
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:49.33-49.54|./core/cpu.sv:195.12-217.6
	value<1> p_cpu__unit_2e_memory__unit_2e_write__req__outstanding;
	// \hdlname: cpu_unit memory_unit read_req_outstanding
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:49.11-49.31|./core/cpu.sv:195.12-217.6
	value<1> p_cpu__unit_2e_memory__unit_2e_read__req__outstanding;
	// \hdlname: cpu_unit memory_unit state_reg
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:22.27-22.36|./core/cpu.sv:195.12-217.6
	value<3> p_cpu__unit_2e_memory__unit_2e_state__reg;
	// \hdlname: cpu_unit memory_unit mem_read_res_valid
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:11.21-11.39|./core/cpu.sv:195.12-217.6
	value<1> p_cpu__unit_2e_memory__unit_2e_mem__read__res__valid;
	// \hdlname: cpu_unit memory_unit ext_ctrl
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:5.27-5.35|./core/cpu.sv:195.12-217.6
	value<3> p_cpu__unit_2e_memory__unit_2e_ext__ctrl;
	// \src: ./soc_top.sv:43.26-47.27
	value<32> i_or_24__2e__2f_soc__top_2e_sv_3a_43_24_12__Y;
	// \src: ./soc_top.sv:32.27-36.28
	value<1> i_or_24__2e__2f_soc__top_2e_sv_3a_32_24_4__Y;
	// \src: ./soc_top.sv:159.16-159.26
	value<32> p_gpio__inout;
	// \src: ./soc_top.sv:133.18-133.26
	value<32> p_gpo__data;
	// \src: ./soc_top.sv:53.18-53.29
	value<32> p_bus__mo__data;
	// \src: ./soc_top.sv:51.11-51.17
	value<1> p_bus__wr;
	// \src: ./soc_top.sv:50.18-50.26
	value<32> p_bus__addr;
	// \src: ./soc_top.sv:49.11-49.17
	value<1> p_bus__rd;
	// connection
	p_cpu__unit_2e_control__unit_2e_instr = p_cpu__unit_2e_fetch__unit_2e_instr.curr;
	// connection
	p_cpu__unit_2e_control__unit_2e_op = p_cpu__unit_2e_control__unit_2e_instr.slice<6,0>().val();
	// cells $procmux$977 $procmux$978_CMP0
	p_cpu__unit_2e_control__unit_2e_is__jalr = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x67u}) ? value<1>{0x1u} : value<1>{0u});
	// cells $procmux$988 $procmux$989_CMP0
	p_cpu__unit_2e_control__unit_2e_is__jal = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x6fu}) ? value<1>{0x1u} : value<1>{0u});
	// cells $procmux$960 $procmux$961_CMP0
	p_cpu__unit_2e_control__unit_2e_is__load = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x3u}) ? value<1>{0x1u} : value<1>{0u});
	// cells $procmux$951 $procmux$952_CMP0
	p_cpu__unit_2e_control__unit_2e_is__store = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x23u}) ? value<1>{0x1u} : value<1>{0u});
	// connection
	p_cpu__unit_2e_memory__unit_2e_state__reg = p_cpu__unit_2e_state__reg.curr;
	// connection
	p_cpu__unit_2e_mem__write = p_cpu__unit_2e_control__unit_2e_is__store;
	// cells $procmux$765 $procmux$762 $procmux$759 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:127$307
	p_cpu__unit_2e_wb__src = (p_cpu__unit_2e_control__unit_2e_is__load ? value<2>{0x1u} : (p_cpu__unit_2e_control__unit_2e_is__load ? value<2>{0u} : (logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__jal, p_cpu__unit_2e_control__unit_2e_is__jalr) ? value<2>{0u} : value<2>{0x2u})));
	// cells $flatten\cpu_unit.\memory_unit.$logic_and$./core/memory.sv:55$357 $flatten\cpu_unit.\memory_unit.$eq$./core/memory.sv:55$356
	p_cpu__unit_2e_memory__unit_2e_write__req__outstanding = logic_and<1>(p_cpu__unit_2e_mem__write, eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_state__reg, value<3>{0x3u}));
	// cells $flatten\cpu_unit.\memory_unit.$logic_and$./core/memory.sv:52$355 $flatten\cpu_unit.\memory_unit.$eq$./core/memory.sv:52$354 $flatten\cpu_unit.\memory_unit.$eq$./core/memory.sv:52$353
	p_cpu__unit_2e_memory__unit_2e_read__req__outstanding = logic_and<1>(eq_uu<1>(p_cpu__unit_2e_wb__src, value<2>{0x1u}), eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_state__reg, value<3>{0x3u}));
	// connection
	p_cpu__unit_2e_fetch__unit_2e_state__reg = p_cpu__unit_2e_state__reg.curr;
	// connection
	p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid = p_cpu__unit_2e_memory__unit_2e_write__req__outstanding;
	// connection
	p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid = p_cpu__unit_2e_memory__unit_2e_read__req__outstanding;
	// \src: ./soc_top.sv:55.9-67.6|./core/fetch.sv:41.35-41.64|./core/cpu.sv:102.11-115.6
	// cell $flatten\cpu_unit.\fetch_unit.$eq$./core/fetch.sv:41$315
	p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid = logic_not<1>(p_cpu__unit_2e_fetch__unit_2e_state__reg);
	// cells $procmux$727 $procmux$691 $procmux$688
	i_procmux_24_727__Y = (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0x1u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<1>{0x1u} : value<1>{0u})));
	// cells $procmux$745 $procmux$715 $procmux$712 $procmux$667 $procmux$664 $procmux$661
	i_procmux_24_745__Y = (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid ? value<1>{0x1u} : value<1>{0u}))))));
	// cells $procmux$724 $procmux$685 $procmux$682 $procmux$649 $procmux$646 $procmux$643
	i_procmux_24_724__Y = (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? p_cpu__unit_2e_fetch__unit_2e_pc__reg.curr : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? p_cpu__unit_2e_alu__result__reg.curr : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid ? p_cpu__unit_2e_alu__result__reg.curr : value<32>{0u}))))));
	// connection
	p_bus__rd = i_procmux_24_727__Y;
	// connection
	p_gpio__unit_2e_i__rd = p_bus__rd;
	// connection
	p_bus__wr = i_procmux_24_745__Y;
	// connection
	p_bus__addr = i_procmux_24_724__Y;
	// connection
	p_gpio__unit_2e_i__wr = p_bus__wr;
	// connection
	p_gpio__unit_2e_i__addr = p_bus__addr;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:47.13-47.32
	// cell $flatten\gpio_unit.$eq$./sys/gpio/gpio.sv:47$24
	i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y = eq_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x300cu});
	// cells $procmux$943 $procmux$944_CMP0
	p_cpu__unit_2e_control__unit_2e_is__iarith = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x13u}) ? value<1>{0x1u} : value<1>{0u});
	// cells $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:65$286 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:65$285
	i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_65_24_286__Y = logic_or<1>(logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__iarith, p_cpu__unit_2e_control__unit_2e_is__jalr), p_cpu__unit_2e_control__unit_2e_is__load);
	// connection
	p_cpu__unit_2e_control__unit_2e_funct3 = p_cpu__unit_2e_control__unit_2e_instr.slice<14,12>().val();
	// cells $procmux$967 $procmux$968_CMP0
	p_cpu__unit_2e_control__unit_2e_is__branch = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x63u}) ? value<1>{0x1u} : value<1>{0u});
	// cells $procmux$1000 $procmux$1001_CMP0
	i_procmux_24_1000__Y = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x17u}) ? value<1>{0x1u} : value<1>{0u});
	// \full_case: 1
	// \src: ./core/control.sv:100.13-105.20
	// cell $procmux$795_CMP0
	i_procmux_24_795__CMP.slice<0>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x6u});
	// \full_case: 1
	// \src: ./core/control.sv:100.13-105.20
	// cell $procmux$795_CMP1
	i_procmux_24_795__CMP.slice<1>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x7u});
	// \full_case: 1
	// \src: ./core/control.sv:100.13-105.20
	// cell $procmux$796_CMP0
	i_procmux_24_796__CMP.slice<0>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x4u});
	// \full_case: 1
	// \src: ./core/control.sv:100.13-105.20
	// cell $procmux$796_CMP1
	i_procmux_24_796__CMP.slice<1>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x5u});
	// connection
	p_gpo__unit_2e_i__addr = p_bus__addr;
	// cells $flatten\gpio_unit.$logic_and$./sys/gpio/gpio.sv:59$33 $flatten\gpio_unit.$eq$./sys/gpio/gpio.sv:59$32
	i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_59_24_33__Y = logic_and<1>(eq_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x3008u}), p_gpio__unit_2e_i__rd);
	// connection
	p_cpu__unit_2e_control__unit_2e_is__auipc = i_procmux_24_1000__Y;
	// cells $procmux$912 $procmux$913_CMP0
	p_cpu__unit_2e_control__unit_2e_is__lui = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x37u}) ? value<1>{0x1u} : value<1>{0u});
	// connection
	p_ram__unit_2e_i__addr = p_bus__addr;
	// cells $procmux$467 $procmux$461 $procmux$459 $procmux$434 $procmux$432 $procmux$429
	i_flatten_5c_gpio__unit_2e__24_1_5c_rd__valid_5b_0_3a_0_5d_ = (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (p_gpio__unit_2e_i__wr ? value<1>{0u} : (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (p_gpio__unit_2e_i__wr ? value<1>{0u} : (p_gpio__unit_2e_i__rd ? value<1>{0x1u} : value<1>{0u})) : value<1>{0u})) : value<1>{0u}) : value<1>{0u});
	// cells $flatten\gpio_unit.$logic_and$./sys/gpio/gpio.sv:57$29 $flatten\gpio_unit.$eq$./sys/gpio/gpio.sv:57$28
	i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y = logic_and<1>(eq_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x3004u}), p_gpio__unit_2e_i__wr);
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:62.13-62.31|./core/cpu.sv:131.13-142.6
	// cell $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:62$284
	i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y = logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__lui, p_cpu__unit_2e_control__unit_2e_is__auipc);
	// cells $flatten\gpio_unit.$logic_and$./sys/gpio/gpio.sv:55$26 $flatten\gpio_unit.$eq$./sys/gpio/gpio.sv:55$25
	i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y = logic_and<1>(eq_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x3000u}), p_gpio__unit_2e_i__wr);
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:97.18-97.37|./core/cpu.sv:131.13-142.6
	// cell $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:97$304
	i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_97_24_304__Y = logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__store, p_cpu__unit_2e_control__unit_2e_is__load);
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:22.25-22.45
	// cell $flatten\timer_unit.$eq$./sys/timer.sv:22$90
	p_timer__unit_2e_timer__addr = eq_uu<1>(p_bus__addr, value<32>{0x4000u});
	// cells $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:96$303 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:96$302
	i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y = logic_or<1>(logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__lui, p_cpu__unit_2e_control__unit_2e_is__auipc), p_cpu__unit_2e_control__unit_2e_is__jalr);
	// cells $procmux$937 $procmux$938_CMP0
	p_cpu__unit_2e_control__unit_2e_is__rarith = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0x33u}) ? value<1>{0x1u} : value<1>{0u});
	// cells $flatten\ram_unit.$logic_and$./sys/ram.sv:28$110 $flatten\ram_unit.$lt$./sys/ram.sv:28$109 $flatten\ram_unit.$ge$./sys/ram.sv:28$108
	p_ram__unit_2e_ram__addr = logic_and<1>(ge_uu<1>(p_ram__unit_2e_i__addr, value<32>{0u}), lt_uu<1>(p_ram__unit_2e_i__addr, value<32>{0x2000u}));
	// \src: ./soc_top.sv:91.11-103.6|./sys/timer.sv:43.13-43.31
	// cell $flatten\timer_unit.$logic_and$./sys/timer.sv:43$93
	i_flatten_5c_timer__unit_2e__24_logic__and_24__2e__2f_sys_2f_timer_2e_sv_3a_43_24_93__Y = logic_and<1>(p_timer__unit_2e_timer__addr, p_bus__rd);
	// connection
	p_gpi__unit_2e_i__addr = p_bus__addr;
	// \src: ./soc_top.sv:55.9-67.6|./core/control.sv:95.13-95.35|./core/cpu.sv:131.13-142.6
	// cell $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:95$301
	i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y = logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__rarith, p_cpu__unit_2e_control__unit_2e_is__iarith);
	// cells $flatten\gpio_unit.$logic_and$./sys/gpio/gpio.sv:21$19 $flatten\gpio_unit.$lt$./sys/gpio/gpio.sv:21$18 $flatten\gpio_unit.$ge$./sys/gpio/gpio.sv:21$17
	p_gpio__unit_2e_gpio__addr = logic_and<1>(ge_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x3000u}), lt_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x3010u}));
	// cells $flatten\gpi_unit.$logic_and$./sys/gpio/gpi.sv:32$379 $flatten\gpi_unit.$eq$./sys/gpio/gpi.sv:32$378
	i_flatten_5c_gpi__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpi_2e_sv_3a_32_24_379__Y = logic_and<1>(eq_uu<1>(p_gpi__unit_2e_i__addr, value<32>{0x2008u}), p_bus__rd);
	// cells $flatten\gpi_unit.$logic_and$./sys/gpio/gpi.sv:24$376 $flatten\gpi_unit.$lt$./sys/gpio/gpi.sv:24$375 $flatten\gpi_unit.$ge$./sys/gpio/gpi.sv:24$374
	p_gpi__unit_2e_gpi__addr = logic_and<1>(ge_uu<1>(p_gpi__unit_2e_i__addr, value<32>{0x2008u}), lt_uu<1>(p_gpi__unit_2e_i__addr, value<32>{0x200cu}));
	// connection
	p_cpu__unit_2e_ext__ctrl = p_cpu__unit_2e_control__unit_2e_funct3;
	// connection
	p_cpu__unit_2e_memory__unit_2e_ext__ctrl = p_cpu__unit_2e_ext__ctrl;
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$68
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_68__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<31>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<31>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$67
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_67__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<30>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<30>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$66
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_66__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<29>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<29>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$65
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_65__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<28>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<28>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$64
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_64__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<27>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<27>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$63
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_63__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<26>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<26>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$62
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_62__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<25>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<25>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$61
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_61__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<24>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<24>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$60
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_60__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<23>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<23>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$59
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_59__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<22>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<22>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$58
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_58__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<21>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<21>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$57
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_57__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<20>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<20>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$56
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_56__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<19>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<19>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$55
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_55__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<18>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<18>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$54
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_54__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<17>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<17>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$53
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_53__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<16>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<16>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$52
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_52__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<15>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<15>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$51
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_51__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<14>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<14>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$50
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_50__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<13>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<13>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$49
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_49__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<12>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<12>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$48
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_48__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<11>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<11>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$47
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_47__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<10>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<10>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$46
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_46__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<9>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<9>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$45
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_45__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<8>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<8>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$44
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_44__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<7>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<7>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$43
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_43__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<6>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<6>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$42
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_42__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<5>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<5>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$41
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_41__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<4>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<4>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$40
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_40__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<3>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<3>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$39
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_39__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<2>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<2>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$38
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_38__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<1>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<1>().val());
	// \src: ./soc_top.sv:161.10-175.6|./sys/gpio/gpio.sv:72.36-72.72
	// cell $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:72$37
	i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_37__Y = (p_gpio__unit_2e_gpio__fsel__reg.curr.slice<0>().val() ? value<1>{0u} : p_gpio__unit_2e_gpo__reg.curr.slice<0>().val());
	// cells $procmux$464 $procmux$455 $procmux$453 $procmux$443 $procmux$441 $procmux$438
	i_flatten_5c_gpio__unit_2e__24_1_5c_rd__data_5b_31_3a_0_5d_ = (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (p_gpio__unit_2e_i__wr ? value<32>{0u} : (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (p_gpio__unit_2e_i__wr ? value<32>{0u} : (p_gpio__unit_2e_i__rd ? p_gpio__unit_2e_gpio__fsel__reg.curr : value<32>{0u})) : value<32>{0u})) : value<32>{0u}) : value<32>{0u});
	// cells $or$./soc_top.sv:32$4 $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:65$34 $procmux$425 $procmux$416 $procmux$413 $procmux$389 $procmux$386 $procmux$383 $or$./soc_top.sv:32$3 $flatten\gpi_unit.$ternary$./sys/gpio/gpi.sv:40$381 $procmux$486 $or$./soc_top.sv:32$1 $flatten\timer_unit.$ternary$./sys/timer.sv:53$101 $procmux$492 $flatten\ram_unit.$ternary$./sys/ram.sv:83$166
	i_or_24__2e__2f_soc__top_2e_sv_3a_32_24_4__Y = or_uu<1>(or_uu<1>(or_uu<1>((p_ram__unit_2e_ram__addr ? p_ram__unit_2e_rd__valid.curr : value<1>{0u}), (p_timer__unit_2e_timer__addr ? (i_flatten_5c_timer__unit_2e__24_logic__and_24__2e__2f_sys_2f_timer_2e_sv_3a_43_24_93__Y ? value<1>{0x1u} : value<1>{0u}) : value<1>{0u})), (p_gpi__unit_2e_gpi__addr ? (i_flatten_5c_gpi__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpi_2e_sv_3a_32_24_379__Y ? value<1>{0x1u} : value<1>{0u}) : value<1>{0u})), (p_gpio__unit_2e_gpio__addr ? (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? i_flatten_5c_gpio__unit_2e__24_1_5c_rd__valid_5b_0_3a_0_5d_ : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? value<1>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y ? i_flatten_5c_gpio__unit_2e__24_1_5c_rd__valid_5b_0_3a_0_5d_ : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? value<1>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y ? value<1>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_59_24_33__Y ? value<1>{0x1u} : i_flatten_5c_gpio__unit_2e__24_1_5c_rd__valid_5b_0_3a_0_5d_)))))) : value<1>{0u}));
	// cells $procmux$899 $procmux$896 $procmux$893 $procmux$890 $procmux$887 $procmux$884 $procmux$881 $procmux$878 $procmux$875 $procmux$872 $procmux$869 $procmux$866 $procmux$863 $procmux$860 $procmux$857
	p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl = (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y ? value<3>{0x4u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0x3u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__branch ? value<3>{0x2u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__branch ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_65_24_286__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_62_24_284__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__branch ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_65_24_286__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__store ? value<3>{0x1u} : value<3>{0x7u})))))))))))))));
	// cells $procmux$736 $procmux$679 $procmux$676
	p_cpu__unit_2e_memory__unit_2e_mem__read__res__valid = (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? i_or_24__2e__2f_soc__top_2e_sv_3a_32_24_4__Y : value<1>{0u})));
	// \src: ./soc_top.sv:55.9-67.6|./core/memory.sv:61.13-61.55|./core/cpu.sv:195.12-217.6
	// cell $flatten\cpu_unit.\memory_unit.$logic_and$./core/memory.sv:61$359
	i_flatten_5c_cpu__unit_2e__5c_memory__unit_2e__24_logic__and_24__2e__2f_core_2f_memory_2e_sv_3a_61_24_359__Y = logic_and<1>(p_cpu__unit_2e_memory__unit_2e_read__req__outstanding, p_cpu__unit_2e_memory__unit_2e_mem__read__res__valid);
	// \full_case: 1
	// \src: ./core/control.sv:113.13-117.20
	// cell $procmux$785_CMP0
	i_procmux_24_785__CMP.slice<0>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x1u});
	// \full_case: 1
	// \src: ./core/control.sv:113.13-117.20
	// cell $procmux$785_CMP1
	i_procmux_24_785__CMP.slice<1>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x4u});
	// \full_case: 1
	// \src: ./core/control.sv:113.13-117.20
	// cell $procmux$785_CMP2
	i_procmux_24_785__CMP.slice<2>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x6u});
	// \full_case: 1
	// \src: ./core/control.sv:113.13-117.20
	// cell $procmux$786_CMP0
	i_procmux_24_786__CMP.slice<0>() = logic_not<1>(p_cpu__unit_2e_control__unit_2e_funct3);
	// \full_case: 1
	// \src: ./core/control.sv:113.13-117.20
	// cell $procmux$786_CMP1
	i_procmux_24_786__CMP.slice<1>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x5u});
	// \full_case: 1
	// \src: ./core/control.sv:113.13-117.20
	// cell $procmux$786_CMP2
	i_procmux_24_786__CMP.slice<2>() = eq_uu<1>(p_cpu__unit_2e_control__unit_2e_funct3, value<3>{0x7u});
	// cells $procmux$739 $procmux$703 $procmux$700 $procmux$640 $procmux$637 $procmux$634
	i_procmux_24_739__Y = (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid ? p_cpu__unit_2e_store__data__reg.curr : value<32>{0u}))))));
	// connection
	p_cpu__unit_2e_execute__unit_2e_rd2 = p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr;
	// connection
	p_cpu__unit_2e_bus__unit_2e_o__bus__data = i_procmux_24_739__Y;
	// connection
	p_cpu__unit_2e_pc__reg = p_cpu__unit_2e_fetch__unit_2e_pc__reg.curr;
	// cells $procmux$752 $procmux$753_CMP0 $procmux$754_CMP0 $procmux$755_CMP0 $procmux$756_CMP0 $procmux$757_CMP0
	p_cpu__unit_2e_imm = (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl, value<3>{0x4u}) ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31,12>().concat(value<12>{0u}).val() : (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl, value<3>{0x3u}) ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31>().val().repeat<12>().concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<19,12>()).concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<20>()).concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<30,21>()).concat(value<1>{0u}).val() : (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl, value<3>{0x2u}) ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31>().val().repeat<20>().concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<7>()).concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<30,25>()).concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<11,8>()).concat(value<1>{0u}).val() : (eq_uu<1>(p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl, value<3>{0x1u}) ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31>().val().repeat<20>().concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31,25>()).concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<11,7>()).val() : (logic_not<1>(p_cpu__unit_2e_control__unit_2e_extract__imm__unit_2e_ctrl) ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31>().val().repeat<20>().concat(p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<31,20>()).val() : value<32>{0u})))));
	// cells $or$./soc_top.sv:43$12 $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:67$36 $procmux$422 $procmux$410 $procmux$407 $procmux$398 $procmux$395 $procmux$392 $or$./soc_top.sv:43$11 $flatten\gpi_unit.$ternary$./sys/gpio/gpi.sv:39$380 $procmux$489 $not$./soc_top.sv:111$13 $or$./soc_top.sv:43$10 $or$./soc_top.sv:43$9 $flatten\timer_unit.$ternary$./sys/timer.sv:54$102 $procmux$495 $flatten\ram_unit.$ternary$./sys/ram.sv:52$116
	i_or_24__2e__2f_soc__top_2e_sv_3a_43_24_12__Y = or_uu<32>(or_uu<32>(or_uu<32>(or_uu<32>((p_ram__unit_2e_ram__addr ? p_ram__unit_2e_read__word.curr : value<32>{0u}), (p_timer__unit_2e_timer__addr ? (i_flatten_5c_timer__unit_2e__24_logic__and_24__2e__2f_sys_2f_timer_2e_sv_3a_43_24_93__Y ? p_timer__unit_2e_timer__reg.curr : value<32>{0u}) : value<32>{0u})), value<32>{0u}), (p_gpi__unit_2e_gpi__addr ? (i_flatten_5c_gpi__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpi_2e_sv_3a_32_24_379__Y ? value<31>{0u}.concat(not_u<1>(p_usr__btn)).val() : value<32>{0u}) : value<32>{0u})), (p_gpio__unit_2e_gpio__addr ? (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? i_flatten_5c_gpio__unit_2e__24_1_5c_rd__data_5b_31_3a_0_5d_ : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? value<32>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y ? i_flatten_5c_gpio__unit_2e__24_1_5c_rd__data_5b_31_3a_0_5d_ : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? value<32>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y ? value<32>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_59_24_33__Y ? i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_68__Y.concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_67__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_66__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_65__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_64__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_63__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_62__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_61__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_60__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_59__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_58__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_57__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_56__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_55__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_54__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_53__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_52__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_51__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_50__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_49__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_48__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_47__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_46__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_45__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_44__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_43__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_42__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_41__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_40__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_39__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_38__Y).concat(i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_37__Y).val() : i_flatten_5c_gpio__unit_2e__24_1_5c_rd__data_5b_31_3a_0_5d_)))))) : value<32>{0u}));
	// connection
	p_bus__mo__data = p_cpu__unit_2e_bus__unit_2e_o__bus__data;
	// cells $flatten\cpu_unit.\fetch_unit.$ternary$./core/fetch.sv:45$316 $procmux$721
	i_flatten_5c_cpu__unit_2e__5c_fetch__unit_2e__24_ternary_24__2e__2f_core_2f_fetch_2e_sv_3a_45_24_316__Y = ((p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? i_or_24__2e__2f_soc__top_2e_sv_3a_32_24_4__Y : value<1>{0u}) ? value<3>{0x1u} : value<3>{0u});
	// connection
	p_gpio__unit_2e_i__data = p_bus__mo__data;
	// connection
	p_gpo__unit_2e_i__data = p_bus__mo__data;
	// connection
	p_gpo__unit_2e_i__wr = p_bus__wr;
	// cells $flatten\ram_unit.$shl$./sys/ram.sv:59$119 $procmux$742 $procmux$709 $procmux$706 $procmux$658 $procmux$655 $procmux$652 $procmux$1046 $procmux$1047_CMP0 $procmux$1048_CMP0 $procmux$1049_CMP0 $procmux$1050_CMP0
	p_ram__unit_2e_write__mask = shl_uu<4>((p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<4>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<4>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<4>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<4>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<4>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid ? (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_ext__ctrl, value<3>{0x5u}) ? value<4>{0x3u} : (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_ext__ctrl, value<3>{0x4u}) ? value<4>{0x1u} : (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_ext__ctrl, value<3>{0x1u}) ? value<4>{0x3u} : (logic_not<1>(p_cpu__unit_2e_memory__unit_2e_ext__ctrl) ? value<4>{0x1u} : value<4>{0xfu})))) : value<4>{0u})))))), p_ram__unit_2e_i__addr.slice<1,0>().val());
	// cells $procmux$608 $procmux$609_CMP0 $procmux$610_CMP0 $procmux$611_CMP0 $procmux$612_CMP0 $procmux$613_CMP0 $procmux$1043 $procmux$1040 $procmux$1037 $flatten\cpu_unit.\memory_unit.$logic_and$./core/memory.sv:64$360 $procmux$730 $procmux$697 $procmux$694 $procmux$631 $procmux$628 $procmux$625 $or$./soc_top.sv:38$8 $flatten\gpio_unit.$ternary$./sys/gpio/gpio.sv:66$35 $or$./soc_top.sv:38$6 $flatten\gpo_unit.$ternary$./sys/gpio/gpo.sv:57$89 $flatten\gpo_unit.$logic_and$./sys/gpio/gpo.sv:24$79 $flatten\gpo_unit.$lt$./sys/gpio/gpo.sv:24$78 $flatten\gpo_unit.$ge$./sys/gpio/gpo.sv:24$77 $flatten\ram_unit.$ternary$./sys/ram.sv:84$167 $flatten\cpu_unit.\execute_unit.$ternary$./core/execute.sv:50$372 $flatten\cpu_unit.\execute_unit.$logic_or$./core/execute.sv:50$371 $flatten\cpu_unit.\execute_unit.$eq$./core/execute.sv:50$370 $flatten\cpu_unit.\decode_unit.$ternary$./core/decode.sv:33$322 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:133$309 $flatten\cpu_unit.\control_unit.$eq$./core/control.sv:133$308 $procmux$932 $procmux$933_CMP0
	i_procmux_24_608__Y = (eq_uu<1>(p_cpu__unit_2e_state__reg.curr, value<3>{0x4u}) ? value<3>{0u} : (eq_uu<1>(p_cpu__unit_2e_state__reg.curr, value<3>{0x3u}) ? (i_flatten_5c_cpu__unit_2e__5c_memory__unit_2e__24_logic__and_24__2e__2f_core_2f_memory_2e_sv_3a_61_24_359__Y ? value<3>{0x4u} : (i_flatten_5c_cpu__unit_2e__5c_memory__unit_2e__24_logic__and_24__2e__2f_core_2f_memory_2e_sv_3a_61_24_359__Y ? value<3>{0u} : (logic_and<1>(p_cpu__unit_2e_memory__unit_2e_write__req__outstanding, (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? value<1>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__write__req__valid ? or_uu<1>(or_uu<1>((p_ram__unit_2e_ram__addr ? p_ram__unit_2e_wr__valid.curr : value<1>{0u}), (logic_and<1>(ge_uu<1>(p_gpo__unit_2e_i__addr, value<32>{0x2000u}), lt_uu<1>(p_gpo__unit_2e_i__addr, value<32>{0x2008u})) ? p_gpo__unit_2e_wr__valid__reg.curr : value<1>{0u})), (p_gpio__unit_2e_gpio__addr ? p_gpio__unit_2e_wr__valid.curr : value<1>{0u})) : value<1>{0u}))))))) ? value<3>{0u} : value<3>{0x3u}))) : (eq_uu<1>(p_cpu__unit_2e_state__reg.curr, value<3>{0x2u}) ? (logic_or<1>(p_cpu__unit_2e_mem__write, eq_uu<1>(p_cpu__unit_2e_wb__src, value<2>{0x1u})) ? value<3>{0x3u} : value<3>{0x4u}) : (eq_uu<1>(p_cpu__unit_2e_state__reg.curr, value<3>{0x1u}) ? (logic_or<1>((eq_uu<1>(p_cpu__unit_2e_control__unit_2e_op, value<7>{0xfu}) ? value<1>{0x1u} : value<1>{0u}), eq_uu<1>(p_cpu__unit_2e_control__unit_2e_instr, value<32>{0x13u})) ? value<3>{0u} : value<3>{0x2u}) : (logic_not<1>(p_cpu__unit_2e_state__reg.curr) ? i_flatten_5c_cpu__unit_2e__5c_fetch__unit_2e__24_ternary_24__2e__2f_core_2f_fetch_2e_sv_3a_45_24_316__Y : value<3>{0x7u})))));
	// \src: ./soc_top.sv:55.9-67.6|./core/execute.sv:29.20-29.52|./core/cpu.sv:166.13-180.6
	// cell $flatten\cpu_unit.\execute_unit.$ternary$./core/execute.sv:29$365
	p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a = (i_procmux_24_1000__Y ? p_cpu__unit_2e_pc__reg : p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r0__data.curr);
	// cells $flatten\cpu_unit.\execute_unit.$ternary$./core/execute.sv:30$369 $flatten\cpu_unit.\control_unit.$ternary$./core/control.sv:90$294 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:90$293 $flatten\cpu_unit.\execute_unit.$ternary$./core/execute.sv:30$368 $flatten\cpu_unit.\control_unit.$ternary$./core/control.sv:91$299 $flatten\cpu_unit.\control_unit.$logic_and$./core/control.sv:91$298 $flatten\cpu_unit.\control_unit.$logic_and$./core/control.sv:91$296 $flatten\cpu_unit.\control_unit.$eq$./core/control.sv:91$295 $flatten\cpu_unit.\execute_unit.$neg$./core/execute.sv:30$367
	p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b = ((logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__branch, p_cpu__unit_2e_control__unit_2e_is__rarith) ? value<1>{0x1u} : value<1>{0u}) ? ((logic_and<1>(logic_and<1>(p_cpu__unit_2e_control__unit_2e_is__rarith, logic_not<1>(p_cpu__unit_2e_control__unit_2e_funct3)), p_cpu__unit_2e_control__unit_2e_instr.slice<30>().val()) ? value<1>{0x1u} : value<1>{0u}) ? neg_u<32>(p_cpu__unit_2e_execute__unit_2e_rd2) : p_cpu__unit_2e_execute__unit_2e_rd2) : p_cpu__unit_2e_imm);
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:63.13-63.29
	// cell $flatten\ram_unit.$logic_and$./sys/ram.sv:63$133
	i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y = logic_and<1>(p_bus__wr, p_ram__unit_2e_ram__addr);
	// connection
	p_ram__unit_2e_write__word = i_procmux_24_739__Y;
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:222.23-222.37
	// cell $flatten\cpu_unit.$add$./core/cpu.sv:222$327
	i_flatten_5c_cpu__unit_2e__24_add_24__2e__2f_core_2f_cpu_2e_sv_3a_222_24_327__Y = add_uu<32>(p_cpu__unit_2e_pc__reg, value<32>{0x4u});
	// cells $flatten\gpo_unit.$logic_and$./sys/gpio/gpo.sv:41$83 $flatten\gpo_unit.$eq$./sys/gpio/gpo.sv:41$82
	i_flatten_5c_gpo__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpo_2e_sv_3a_41_24_83__Y = logic_and<1>(eq_uu<1>(p_gpo__unit_2e_i__addr, value<32>{0x2000u}), p_gpo__unit_2e_i__wr);
	// connection
	p_cpu__unit_2e_pc__plus4 = i_flatten_5c_cpu__unit_2e__24_add_24__2e__2f_core_2f_cpu_2e_sv_3a_222_24_327__Y;
	// cells $procmux$790 $procmux$787 $procmux$784 $procmux$785_ANY $procmux$786_ANY $procmux$780 $procmux$777 $procmux$774 $procmux$771 $procmux$768
	p_cpu__unit_2e_next__pc = (p_cpu__unit_2e_control__unit_2e_is__branch ? (p_cpu__unit_2e_control__unit_2e_is__branch ? (reduce_or<1>(i_procmux_24_785__CMP) ? value<2>{0x3u} : (reduce_or<1>(i_procmux_24_786__CMP) ? value<2>{0x2u} : value<2>{0u})) : value<2>{0u}) : (p_cpu__unit_2e_control__unit_2e_is__branch ? value<2>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<2>{0x2u} : (p_cpu__unit_2e_control__unit_2e_is__branch ? value<2>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<2>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jalr ? value<2>{0u} : value<2>{0x1u}))))));
	// connection
	p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl = p_cpu__unit_2e_memory__unit_2e_ext__ctrl;
	// cells $procmux$854 $procmux$851 $procmux$848 $procmux$845 $procmux$842 $procmux$839 $procmux$836 $procmux$833 $procmux$830 $procmux$827 $procmux$824 $procmux$821 $procmux$818 $procmux$815 $procmux$812 $procmux$809 $procmux$806 $procmux$803 $procmux$800 $procmux$797 $procmux$794 $procmux$795_ANY $procmux$796_ANY
	p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op = (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<14,12>().val() : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_97_24_304__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_97_24_304__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0x7u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_97_24_304__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__branch ? (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_95_24_301__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_96_24_303__Y ? value<3>{0u} : (i_flatten_5c_cpu__unit_2e__5c_control__unit_2e__24_logic__or_24__2e__2f_core_2f_control_2e_sv_3a_97_24_304__Y ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__jal ? value<3>{0u} : (p_cpu__unit_2e_control__unit_2e_is__branch ? (reduce_or<1>(i_procmux_24_795__CMP) ? value<3>{0x3u} : (reduce_or<1>(i_procmux_24_796__CMP) ? value<3>{0x2u} : value<3>{0x4u})) : value<3>{0u}))))) : value<3>{0x7u})))))))))))))));
	// connection
	p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl = p_cpu__unit_2e_ext__ctrl;
	// cells $flatten\cpu_unit.\decode_unit.\reg_file_unit.$logic_and$./core/reg_file.sv:25$210 $flatten\cpu_unit.\decode_unit.\reg_file_unit.$ne$./core/reg_file.sv:25$209 $flatten\cpu_unit.$logic_and$./core/cpu.sv:226$329 $flatten\cpu_unit.\control_unit.$logic_not$./core/control.sv:86$291 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:86$290 $flatten\cpu_unit.$eq$./core/cpu.sv:226$328
	i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_logic__and_24__2e__2f_core_2f_reg__file_2e_sv_3a_25_24_210__Y = logic_and<1>(logic_and<1>(eq_uu<1>(p_cpu__unit_2e_state__reg.curr, value<3>{0x4u}), logic_not<1>(logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__store, p_cpu__unit_2e_control__unit_2e_is__branch))), reduce_bool<1>(p_cpu__unit_2e_control__unit_2e_instr.slice<11,7>().val()));
	// cells $procmux$733 $procmux$673 $procmux$670
	p_cpu__unit_2e_bus__unit_2e_mem__read__res__data = (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? value<32>{0u} : (p_cpu__unit_2e_bus__unit_2e_mem__read__req__valid ? i_or_24__2e__2f_soc__top_2e_sv_3a_43_24_12__Y : value<32>{0u})));
	// connection
	p_cpu__unit_2e_clk = p_clk;
	// \full_case: 1
	// \src: ./core/reg_file.sv:25.9-26.36
	// cell $procmux$1009
	i_procmux_24_1009__Y = (i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_logic__and_24__2e__2f_core_2f_reg__file_2e_sv_3a_25_24_210__Y ? p_cpu__unit_2e_fetch__unit_2e_instr.curr.slice<11,7>().val() : value<5>{0u});
	// cells $procmux$1006 $procmux$615 $procmux$616_CMP0 $procmux$617_CMP0 $procmux$618_CMP0
	i_procmux_24_1006__Y = (i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_logic__and_24__2e__2f_core_2f_reg__file_2e_sv_3a_25_24_210__Y ? (eq_uu<1>(p_cpu__unit_2e_wb__src, value<2>{0x2u}) ? p_cpu__unit_2e_alu__result__reg.curr : (eq_uu<1>(p_cpu__unit_2e_wb__src, value<2>{0x1u}) ? p_cpu__unit_2e_memory__unit_2e_mem__rdata__reg.curr : (logic_not<1>(p_cpu__unit_2e_wb__src) ? i_flatten_5c_cpu__unit_2e__24_add_24__2e__2f_core_2f_cpu_2e_sv_3a_222_24_327__Y : value<32>{0u}))) : value<32>{0u});
	// \full_case: 1
	// \src: ./core/reg_file.sv:25.9-26.36
	// cell $procmux$1003
	i_procmux_24_1003__Y = (i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_logic__and_24__2e__2f_core_2f_reg__file_2e_sv_3a_25_24_210__Y ? value<32>{0xffffffffu} : value<32>{0u});
	// cells $procmux$603 $procmux$567 $procmux$565
	i_procmux_24_603__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<3>().val() ? value<2>{0u}.concat(i_procmux_24_724__Y.slice<31,2>()).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$600 $procmux$561 $procmux$559
	i_procmux_24_600__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<3>().val() ? p_ram__unit_2e_write__word.slice<31,24>().concat(value<24>{0u}).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$597 $procmux$555 $procmux$553
	i_procmux_24_597__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<3>().val() ? value<32>{0xff000000u} : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$594 $procmux$549 $procmux$547
	i_procmux_24_594__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<2>().val() ? value<2>{0u}.concat(i_procmux_24_724__Y.slice<31,2>()).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$591 $procmux$543 $procmux$541
	i_procmux_24_591__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<2>().val() ? value<8>{0u}.concat(p_ram__unit_2e_write__word.slice<23,16>()).concat(value<16>{0u}).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$588 $procmux$537 $procmux$535
	i_procmux_24_588__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<2>().val() ? value<32>{0xff0000u} : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$585 $procmux$531 $procmux$529
	i_procmux_24_585__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<1>().val() ? value<2>{0u}.concat(i_procmux_24_724__Y.slice<31,2>()).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$582 $procmux$525 $procmux$523
	i_procmux_24_582__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<1>().val() ? value<16>{0u}.concat(p_ram__unit_2e_write__word.slice<15,8>()).concat(value<8>{0u}).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$579 $procmux$519 $procmux$517
	i_procmux_24_579__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<1>().val() ? value<32>{0xff00u} : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$576 $procmux$513 $procmux$511
	i_procmux_24_576__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<0>().val() ? value<2>{0u}.concat(i_procmux_24_724__Y.slice<31,2>()).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$573 $procmux$507 $procmux$505
	i_procmux_24_573__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<0>().val() ? value<24>{0u}.concat(p_ram__unit_2e_write__word.slice<7,0>()).val() : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// cells $procmux$570 $procmux$501 $procmux$499
	i_procmux_24_570__Y = (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (i_flatten_5c_ram__unit_2e__24_logic__and_24__2e__2f_sys_2f_ram_2e_sv_3a_63_24_133__Y ? (p_ram__unit_2e_write__mask.slice<0>().val() ? value<32>{0xffu} : value<32>{0u}) : value<32>{0u}) : value<32>{0u});
	// connection
	p_cpu__unit_2e_rst = p_rst;
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:76.26-76.29
	// memory \ram_unit.mem read port 0
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:76.30-76.60
	// cell $flatten\ram_unit.$ternary$./sys/ram.sv:76$165
	auto tmp_0 = memory_index((p_ram__unit_2e_ram__addr ? value<2>{0u}.concat(i_procmux_24_724__Y.slice<31,2>()).val() : value<32>{0u}), 0, 2048);
	CXXRTL_ASSERT(tmp_0.valid && "out of bounds read");
	if(tmp_0.valid) {
		value<32> tmp_1 = memory_p_ram__unit_2e_mem[tmp_0.index];
		i_flatten_5c_ram__unit_2e__24_memrd_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_76_24_163__DATA = tmp_1;
	} else {
		i_flatten_5c_ram__unit_2e__24_memrd_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_76_24_163__DATA = value<32> {};
	}
	// connection
	p_gpio__unit_2e_gpio__inout.slice<31>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_68__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<30>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_67__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<29>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_66__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<28>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_65__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<27>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_64__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<26>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_63__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<25>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_62__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<24>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_61__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<23>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_60__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<22>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_59__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<21>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_58__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<20>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_57__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<19>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_56__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<18>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_55__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<17>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_54__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<16>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_53__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<15>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_52__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<14>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_51__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<13>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_50__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<12>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_49__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<11>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_48__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<10>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_47__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<9>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_46__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<8>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_45__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<7>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_44__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<6>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_43__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<5>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_42__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<4>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_41__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<3>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_40__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<2>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_39__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<1>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_38__Y;
	// connection
	p_gpio__unit_2e_gpio__inout.slice<0>() = i_flatten_5c_gpio__unit_2e__24_ternary_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_72_24_37__Y;
	// connection
	p_gpo__data = p_gpo__unit_2e_gpo__reg.curr;
	// connection
	p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_clk = p_cpu__unit_2e_clk;
	// connection
	p_ram__unit_2e_clk = p_clk;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:28.20-28.24|./core/cpu.sv:148.12-159.6
	// memory \cpu_unit.decode_unit.reg_file_unit.regs read port 1
	// cells $flatten\cpu_unit.\control_unit.$ternary$./core/control.sv:77$289 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:77$288 $flatten\cpu_unit.\control_unit.$logic_or$./core/control.sv:77$287
	auto tmp_2 = memory_index((logic_or<1>(logic_or<1>(p_cpu__unit_2e_control__unit_2e_is__jal, p_cpu__unit_2e_control__unit_2e_is__lui), p_cpu__unit_2e_control__unit_2e_is__auipc) ? value<5>{0u} : p_cpu__unit_2e_control__unit_2e_instr.slice<19,15>().val()), 0, 32);
	CXXRTL_ASSERT(tmp_2.valid && "out of bounds read");
	if(tmp_2.valid) {
		value<32> tmp_3 = memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs[tmp_2.index];
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_28_24_214__DATA = tmp_3;
	} else {
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_28_24_214__DATA = value<32> {};
	}
	// connection
	p_cpu__unit_2e_fetch__unit_2e_clk = p_cpu__unit_2e_clk;
	// connection
	p_cpu__unit_2e_fetch__unit_2e_rst = p_cpu__unit_2e_rst;
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:29.20-29.24|./core/cpu.sv:148.12-159.6
	// memory \cpu_unit.decode_unit.reg_file_unit.regs read port 0
	auto tmp_4 = memory_index(p_cpu__unit_2e_control__unit_2e_instr.slice<24,20>().val(), 0, 32);
	CXXRTL_ASSERT(tmp_4.valid && "out of bounds read");
	if(tmp_4.valid) {
		value<32> tmp_5 = memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs[tmp_4.index];
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_29_24_215__DATA = tmp_5;
	} else {
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_29_24_215__DATA = value<32> {};
	}
	// connection
	p_gpo__unit_2e_clk = p_clk;
	// connection
	p_gpo__unit_2e_rst = p_rst;
	// connection
	p_timer__unit_2e_clk = p_clk;
	// connection
	p_timer__unit_2e_rst = p_rst;
	// connection
	p_gpio__unit_2e_clk = p_clk;
	// connection
	p_gpio__unit_2e_rst = p_rst;
	// connection
	p_gpio__inout = p_gpio__unit_2e_gpio__inout;
	// \src: ./core/reg_file.sv:26.13-26.35
	// memory \cpu_unit.decode_unit.reg_file_unit.regs write port 0
	if (posedge_p_clk) {
		auto tmp_6 = memory_index(i_procmux_24_1009__Y, 0, 32);
		CXXRTL_ASSERT(tmp_6.valid && "out of bounds write");
		if (tmp_6.valid) {
			memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs.update(tmp_6.index, i_procmux_24_1006__Y, i_procmux_24_1003__Y, 0);
		}
	}
	// \src: ./sys/ram.sv:66.17-66.61
	// memory \ram_unit.mem write port 0
	if (posedge_p_clk) {
		auto tmp_7 = memory_index(i_procmux_24_603__Y, 0, 2048);
		CXXRTL_ASSERT(tmp_7.valid && "out of bounds write");
		if (tmp_7.valid) {
			memory_p_ram__unit_2e_mem.update(tmp_7.index, i_procmux_24_600__Y, i_procmux_24_597__Y, 0);
		}
	}
	// \src: ./sys/ram.sv:68.17-68.61
	// memory \ram_unit.mem write port 1
	if (posedge_p_clk) {
		auto tmp_8 = memory_index(i_procmux_24_594__Y, 0, 2048);
		CXXRTL_ASSERT(tmp_8.valid && "out of bounds write");
		if (tmp_8.valid) {
			memory_p_ram__unit_2e_mem.update(tmp_8.index, i_procmux_24_591__Y, i_procmux_24_588__Y, 1);
		}
	}
	// \src: ./sys/ram.sv:70.17-70.59
	// memory \ram_unit.mem write port 2
	if (posedge_p_clk) {
		auto tmp_9 = memory_index(i_procmux_24_585__Y, 0, 2048);
		CXXRTL_ASSERT(tmp_9.valid && "out of bounds write");
		if (tmp_9.valid) {
			memory_p_ram__unit_2e_mem.update(tmp_9.index, i_procmux_24_582__Y, i_procmux_24_579__Y, 2);
		}
	}
	// \src: ./sys/ram.sv:72.17-72.57
	// memory \ram_unit.mem write port 3
	if (posedge_p_clk) {
		auto tmp_10 = memory_index(i_procmux_24_576__Y, 0, 2048);
		CXXRTL_ASSERT(tmp_10.valid && "out of bounds write");
		if (tmp_10.valid) {
			memory_p_ram__unit_2e_mem.update(tmp_10.index, i_procmux_24_573__Y, i_procmux_24_570__Y, 3);
		}
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:24.5-30.8|./core/cpu.sv:148.12-159.6
	// cell $procdff$1078
	if (posedge_p_clk) {
		p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r0__data.next = i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_28_24_214__DATA;
	}
	// cells $procdff$1074 $procmux$1018 $procmux$1019_CMP0 $procmux$1020_CMP0 $procmux$1021_CMP0 $procmux$1022_CMP0 $procmux$1023_CMP0 $procmux$1024_CMP0 $procmux$1025_CMP0 $procmux$1026_CMP0 $flatten\cpu_unit.\execute_unit.\alu_unit.$and$./core/alu.sv:24$351 $flatten\cpu_unit.\execute_unit.\alu_unit.$or$./core/alu.sv:23$350 $flatten\cpu_unit.\execute_unit.\alu_unit.$xor$./core/alu.sv:22$349 $flatten\cpu_unit.\execute_unit.\alu_unit.$ternary$./core/alu.sv:21$348 $flatten\cpu_unit.\execute_unit.\alu_unit.$lt$./core/alu.sv:21$347 $flatten\cpu_unit.\execute_unit.\alu_unit.$ternary$./core/alu.sv:20$346 $flatten\cpu_unit.\execute_unit.\alu_unit.$lt$./core/alu.sv:20$345 $flatten\cpu_unit.\execute_unit.\alu_unit.$ternary$./core/alu.sv:19$344 $flatten\cpu_unit.\execute_unit.\alu_unit.$sshr$./core/alu.sv:13$339 $flatten\cpu_unit.\execute_unit.\alu_unit.$shr$./core/alu.sv:19$343 $flatten\cpu_unit.\execute_unit.\alu_unit.$shl$./core/alu.sv:18$342 $flatten\cpu_unit.\execute_unit.\alu_unit.$add$./core/alu.sv:17$341
	if (posedge_p_clk) {
		p_cpu__unit_2e_alu__result__reg.next = (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x7u}) ? and_uu<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b) : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x6u}) ? or_uu<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b) : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x4u}) ? xor_uu<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b) : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x3u}) ? (lt_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b) ? value<32>{0x1u} : value<32>{0u}) : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x2u}) ? (lt_ss<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b) ? value<32>{0x1u} : value<32>{0u}) : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x5u}) ? (p_cpu__unit_2e_control__unit_2e_instr.slice<30>().val() ? sshr_su<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b.slice<4,0>().val()) : shr_uu<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b.slice<4,0>().val())) : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op, value<3>{0x1u}) ? shl_uu<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b.slice<4,0>().val()) : (logic_not<1>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_op) ? add_uu<32>(p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_a, p_cpu__unit_2e_execute__unit_2e_alu__unit_2e_b) : value<32>{0u}))))))));
	}
	if (p_cpu__unit_2e_rst == value<1> {1u}) {
		p_cpu__unit_2e_alu__result__reg.next = value<32>{0u};
	}
	// cells $procdff$1077 $procmux$749 $flatten\cpu_unit.\fetch_unit.$logic_and$./core/fetch.sv:33$314 $flatten\cpu_unit.\fetch_unit.$ne$./core/fetch.sv:33$313 $flatten\cpu_unit.\fetch_unit.$eq$./core/fetch.sv:33$312 $procmux$619 $procmux$620_CMP0 $procmux$621_CMP0 $procmux$622_CMP0 $procmux$623_CMP0 $flatten\cpu_unit.$ternary$./core/cpu.sv:247$337 $flatten\cpu_unit.$eq$./core/cpu.sv:247$335 $flatten\cpu_unit.$add$./core/cpu.sv:247$336 $flatten\cpu_unit.$ternary$./core/cpu.sv:246$334 $flatten\cpu_unit.$eq$./core/cpu.sv:246$332 $flatten\cpu_unit.$add$./core/cpu.sv:246$333
	if (posedge_p_clk) {
		p_cpu__unit_2e_fetch__unit_2e_pc__reg.next = (logic_and<1>(logic_not<1>(i_procmux_24_608__Y), reduce_bool<1>(p_cpu__unit_2e_fetch__unit_2e_state__reg)) ? (eq_uu<1>(p_cpu__unit_2e_next__pc, value<2>{0x3u}) ? (logic_not<1>(p_cpu__unit_2e_alu__result__reg.curr) ? p_cpu__unit_2e_pc__plus4 : add_uu<32>(p_cpu__unit_2e_pc__reg, p_cpu__unit_2e_imm)) : (eq_uu<1>(p_cpu__unit_2e_next__pc, value<2>{0x2u}) ? (logic_not<1>(p_cpu__unit_2e_alu__result__reg.curr) ? add_uu<32>(p_cpu__unit_2e_pc__reg, p_cpu__unit_2e_imm) : p_cpu__unit_2e_pc__plus4) : (eq_uu<1>(p_cpu__unit_2e_next__pc, value<2>{0x1u}) ? i_flatten_5c_cpu__unit_2e__24_add_24__2e__2f_core_2f_cpu_2e_sv_3a_222_24_327__Y : (logic_not<1>(p_cpu__unit_2e_next__pc) ? p_cpu__unit_2e_alu__result__reg.curr.slice<31,1>().concat(value<1>{0u}).val() : value<32>{0u})))) : p_cpu__unit_2e_fetch__unit_2e_pc__reg.curr);
	}
	if (p_cpu__unit_2e_fetch__unit_2e_rst == value<1> {1u}) {
		p_cpu__unit_2e_fetch__unit_2e_pc__reg.next = value<32>{0u};
	}
	// cells $procdff$1083 $procmux$1034 $flatten\cpu_unit.\memory_unit.$logic_and$./core/memory.sv:86$363 $flatten\cpu_unit.\memory_unit.$eq$./core/memory.sv:86$362 $procmux$1028 $procmux$1029_CMP0 $procmux$1030_CMP0 $procmux$1031_CMP0 $procmux$1032_CMP0 $procmux$1033_CMP0
	if (posedge_p_clk) {
		p_cpu__unit_2e_memory__unit_2e_mem__rdata__reg.next = (logic_and<1>(eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_state__reg, value<3>{0x3u}), p_cpu__unit_2e_memory__unit_2e_mem__read__res__valid) ? (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl, value<3>{0x5u}) ? value<16>{0u}.concat(p_cpu__unit_2e_bus__unit_2e_mem__read__res__data.slice<15,0>()).val() : (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl, value<3>{0x4u}) ? value<24>{0u}.concat(p_cpu__unit_2e_bus__unit_2e_mem__read__res__data.slice<7,0>()).val() : (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl, value<3>{0x2u}) ? p_cpu__unit_2e_bus__unit_2e_mem__read__res__data : (eq_uu<1>(p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl, value<3>{0x1u}) ? p_cpu__unit_2e_bus__unit_2e_mem__read__res__data.slice<15>().val().repeat<16>().concat(p_cpu__unit_2e_bus__unit_2e_mem__read__res__data.slice<15,0>()).val() : (logic_not<1>(p_cpu__unit_2e_memory__unit_2e_mem__read__extend__unit_2e_ctrl) ? p_cpu__unit_2e_bus__unit_2e_mem__read__res__data.slice<7>().val().repeat<24>().concat(p_cpu__unit_2e_bus__unit_2e_mem__read__res__data.slice<7,0>()).val() : value<32>{0u}))))) : p_cpu__unit_2e_memory__unit_2e_mem__rdata__reg.curr);
	}
	if (p_cpu__unit_2e_rst == value<1> {1u}) {
		p_cpu__unit_2e_memory__unit_2e_mem__rdata__reg.next = value<32>{0u};
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:24.5-30.8|./core/cpu.sv:148.12-159.6
	// cell $procdff$1079
	if (posedge_p_clk) {
		p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.next = i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memrd_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_29_24_215__DATA;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:24.5-30.8|./core/cpu.sv:148.12-159.6
	// cell $procdff$1080
	if (posedge_p_clk) {
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__ADDR.next = i_procmux_24_1009__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:24.5-30.8|./core/cpu.sv:148.12-159.6
	// cell $procdff$1081
	if (posedge_p_clk) {
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__DATA.next = i_procmux_24_1006__Y;
	}
	// cells $procdff$1075 $procmux$1012 $procmux$1013_CMP0 $procmux$1014_CMP0 $procmux$1015_CMP0 $procmux$1016_CMP0 $procmux$1017_CMP0
	if (posedge_p_clk) {
		p_cpu__unit_2e_store__data__reg.next = (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl, value<3>{0x5u}) ? value<16>{0u}.concat(p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr.slice<15,0>()).val() : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl, value<3>{0x4u}) ? value<24>{0u}.concat(p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr.slice<7,0>()).val() : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl, value<3>{0x2u}) ? p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr : (eq_uu<1>(p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl, value<3>{0x1u}) ? p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr.slice<15>().val().repeat<16>().concat(p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr.slice<15,0>()).val() : (logic_not<1>(p_cpu__unit_2e_execute__unit_2e_reg__extend__unit_2e_ctrl) ? p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr.slice<7>().val().repeat<24>().concat(p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.curr.slice<7,0>()).val() : value<32>{0u})))));
	}
	if (p_cpu__unit_2e_rst == value<1> {1u}) {
		p_cpu__unit_2e_store__data__reg.next = value<32>{0u};
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:55.9-67.6|./core/decode.sv:23.14-28.6|./core/reg_file.sv:24.5-30.8|./core/cpu.sv:148.12-159.6
	// cell $procdff$1082
	if (posedge_p_clk) {
		i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__EN.next = i_procmux_24_1003__Y;
	}
	// cells $procdff$1076 $procmux$747 $flatten\cpu_unit.\fetch_unit.$logic_and$./core/fetch.sv:53$320 $flatten\cpu_unit.\fetch_unit.$eq$./core/fetch.sv:53$319 $flatten\cpu_unit.\fetch_unit.$eq$./core/fetch.sv:53$318 $procmux$718
	if (posedge_p_clk) {
		p_cpu__unit_2e_fetch__unit_2e_instr.next = (logic_and<1>(logic_not<1>(p_cpu__unit_2e_fetch__unit_2e_state__reg), eq_uu<1>(i_flatten_5c_cpu__unit_2e__5c_fetch__unit_2e__24_ternary_24__2e__2f_core_2f_fetch_2e_sv_3a_45_24_316__Y, value<3>{0x1u})) ? (p_cpu__unit_2e_bus__unit_2e_instr__read__req__valid ? i_or_24__2e__2f_soc__top_2e_sv_3a_43_24_12__Y : value<32>{0u}) : p_cpu__unit_2e_fetch__unit_2e_instr.curr);
	}
	if (p_cpu__unit_2e_fetch__unit_2e_rst == value<1> {1u}) {
		p_cpu__unit_2e_fetch__unit_2e_instr.next = value<32>{0u};
	}
	// cells $procdff$1051 $procmux$483 $flatten\gpo_unit.$or$./sys/gpio/gpo.sv:42$84 $procmux$480 $procmux$477 $flatten\gpo_unit.$logic_and$./sys/gpio/gpo.sv:43$86 $flatten\gpo_unit.$eq$./sys/gpio/gpo.sv:43$85 $flatten\gpo_unit.$and$./sys/gpio/gpo.sv:44$88 $flatten\gpo_unit.$not$./sys/gpio/gpo.sv:44$87
	if (posedge_p_clk) {
		p_gpo__unit_2e_gpo__reg.next = (i_flatten_5c_gpo__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpo_2e_sv_3a_41_24_83__Y ? or_uu<32>(p_gpo__unit_2e_gpo__reg.curr, p_gpo__unit_2e_i__data) : (i_flatten_5c_gpo__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpo_2e_sv_3a_41_24_83__Y ? value<32>{0u} : (logic_and<1>(eq_uu<1>(p_gpo__unit_2e_i__addr, value<32>{0x2004u}), p_gpo__unit_2e_i__wr) ? and_uu<32>(p_gpo__unit_2e_gpo__reg.curr, not_u<32>(p_gpo__unit_2e_i__data)) : p_gpo__unit_2e_gpo__reg.curr)));
	}
	if (p_gpo__unit_2e_rst == value<1> {1u}) {
		p_gpo__unit_2e_gpo__reg.next = value<32>{0u};
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:135.9-149.6|./sys/gpio/gpo.sv:29.5-37.8
	// cell $procdff$1052
	if (posedge_p_clk) {
		p_gpo__unit_2e_wr__valid__reg.next = i_procmux_24_745__Y;
	}
	if (p_gpo__unit_2e_rst == value<1> {1u}) {
		p_gpo__unit_2e_wr__valid__reg.next = value<1>{0u};
	}
	// cells $procdff$1053 $flatten\timer_unit.$ternary$./sys/timer.sv:49$96 $flatten\timer_unit.$eq$./sys/timer.sv:49$94 $flatten\timer_unit.$add$./sys/timer.sv:49$95
	if (posedge_p_clk) {
		p_timer__unit_2e_timer__reg.next = (eq_uu<1>(p_timer__unit_2e_tick__reg.curr, value<6>{0x30u}) ? add_uu<32>(p_timer__unit_2e_timer__reg.curr, value<32>{0x1u}) : p_timer__unit_2e_timer__reg.curr);
	}
	if (p_timer__unit_2e_rst == value<1> {1u}) {
		p_timer__unit_2e_timer__reg.next = value<32>{0u};
	}
	// cells $procdff$1054 $flatten\timer_unit.$ternary$./sys/timer.sv:50$99 $flatten\timer_unit.$eq$./sys/timer.sv:50$97 $flatten\timer_unit.$add$./sys/timer.sv:50$98
	if (posedge_p_clk) {
		p_timer__unit_2e_tick__reg.next = (eq_uu<1>(p_timer__unit_2e_tick__reg.curr, value<6>{0x30u}) ? value<32>{0u} : add_uu<32>(p_timer__unit_2e_tick__reg.curr, value<32>{0x1u})).slice<5,0>().val();
	}
	if (p_timer__unit_2e_rst == value<1> {1u}) {
		p_timer__unit_2e_tick__reg.next = value<6>{0u};
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1055
	if (posedge_p_clk) {
		p_ram__unit_2e_wr__valid.next = i_procmux_24_745__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1056
	if (posedge_p_clk) {
		p_ram__unit_2e_rd__valid.next = i_procmux_24_727__Y;
	}
	// cells $procdff$1057 $procmux$605 $flatten\ram_unit.$logic_and$./sys/ram.sv:75$162
	if (posedge_p_clk) {
		p_ram__unit_2e_read__word.next = (logic_and<1>(p_bus__rd, p_ram__unit_2e_ram__addr) ? i_flatten_5c_ram__unit_2e__24_memrd_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_76_24_163__DATA : p_ram__unit_2e_read__word.curr);
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1058
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__ADDR.next = i_procmux_24_603__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1059
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__DATA.next = i_procmux_24_600__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1060
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__EN.next = i_procmux_24_597__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1061
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__ADDR.next = i_procmux_24_594__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1062
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__DATA.next = i_procmux_24_591__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1063
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__EN.next = i_procmux_24_588__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1064
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__ADDR.next = i_procmux_24_585__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1065
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__DATA.next = i_procmux_24_582__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1066
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__EN.next = i_procmux_24_579__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1067
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__ADDR.next = i_procmux_24_576__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1068
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__DATA.next = i_procmux_24_573__Y;
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:73.9-85.6|./sys/ram.sv:62.5-81.8
	// cell $procdff$1069
	if (posedge_p_clk) {
		i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__EN.next = i_procmux_24_570__Y;
	}
	// cells $procdff$1070 $procmux$419 $flatten\gpio_unit.$or$./sys/gpio/gpio.sv:56$27 $procmux$404 $procmux$401 $flatten\gpio_unit.$and$./sys/gpio/gpio.sv:58$31 $flatten\gpio_unit.$not$./sys/gpio/gpio.sv:58$30
	if (posedge_p_clk) {
		p_gpio__unit_2e_gpo__reg.next = (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? or_uu<32>(p_gpio__unit_2e_gpo__reg.curr, p_gpio__unit_2e_i__data) : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_55_24_26__Y ? value<32>{0u} : (i_flatten_5c_gpio__unit_2e__24_logic__and_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_57_24_29__Y ? and_uu<32>(p_gpio__unit_2e_gpo__reg.curr, not_u<32>(p_gpio__unit_2e_i__data)) : p_gpio__unit_2e_gpo__reg.curr)));
	}
	if (p_gpio__unit_2e_rst == value<1> {1u}) {
		p_gpio__unit_2e_gpo__reg.next = value<32>{0u};
	}
	// cells $procdff$1071 $procmux$470 $procmux$449 $procmux$447
	if (posedge_p_clk) {
		p_gpio__unit_2e_gpio__fsel__reg.next = (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (i_flatten_5c_gpio__unit_2e__24_eq_24__2e__2f_sys_2f_gpio_2f_gpio_2e_sv_3a_47_24_24__Y ? (p_gpio__unit_2e_i__wr ? p_cpu__unit_2e_bus__unit_2e_o__bus__data : p_gpio__unit_2e_gpio__fsel__reg.curr) : value<32>{0u}) : p_gpio__unit_2e_gpio__fsel__reg.curr);
	}
	if (p_gpio__unit_2e_rst == value<1> {1u}) {
		p_gpio__unit_2e_gpio__fsel__reg.next = value<32>{0u};
	}
	// cells $procdff$1072 $flatten\gpio_unit.$logic_and$./sys/gpio/gpio.sv:38$22 $flatten\gpio_unit.$ne$./sys/gpio/gpio.sv:38$21
	if (posedge_p_clk) {
		p_gpio__unit_2e_wr__valid.next = logic_and<1>(p_gpio__unit_2e_i__wr, ne_uu<1>(p_gpio__unit_2e_i__addr, value<32>{0x3008u}));
	}
	if (p_gpio__unit_2e_rst == value<1> {1u}) {
		p_gpio__unit_2e_wr__valid.next = value<1>{0u};
	}
	// \always_ff: 1
	// \src: ./soc_top.sv:55.9-67.6|./core/cpu.sv:78.5-84.8
	// cell $procdff$1073
	if (posedge_p_clk) {
		p_cpu__unit_2e_state__reg.next = i_procmux_24_608__Y;
	}
	if (p_cpu__unit_2e_rst == value<1> {1u}) {
		p_cpu__unit_2e_state__reg.next = value<3>{0u};
	}
	// connection
	p_gpio__a3 = p_gpio__inout.slice<12>().val();
	// connection
	p_gpio__a2 = p_gpio__inout.slice<11>().val();
	// connection
	p_gpio__a1 = p_gpio__inout.slice<10>().val();
	// connection
	p_gpio__a0 = p_gpio__inout.slice<9>().val();
	// connection
	p_gpio__13 = p_gpio__inout.slice<8>().val();
	// connection
	p_gpio__12 = p_gpio__inout.slice<7>().val();
	// connection
	p_gpio__11 = p_gpio__inout.slice<6>().val();
	// connection
	p_gpio__10 = p_gpio__inout.slice<5>().val();
	// connection
	p_gpio__9 = p_gpio__inout.slice<4>().val();
	// connection
	p_gpio__6 = p_gpio__inout.slice<3>().val();
	// connection
	p_gpio__5 = p_gpio__inout.slice<2>().val();
	// connection
	p_gpio__1 = p_gpio__inout.slice<1>().val();
	// connection
	p_gpio__0 = p_gpio__inout.slice<0>().val();
	// \src: ./soc_top.sv:153.25-153.37
	// cell $not$./soc_top.sv:153$16
	p_rgb__led0__b = not_u<1>(p_gpo__data.slice<2>().val());
	// \src: ./soc_top.sv:152.25-152.37
	// cell $not$./soc_top.sv:152$15
	p_rgb__led0__g = not_u<1>(p_gpo__data.slice<1>().val());
	// \src: ./soc_top.sv:151.25-151.37
	// cell $not$./soc_top.sv:151$14
	p_rgb__led0__r = not_u<1>(p_gpo__data.slice<0>().val());
	return converged;
}

bool p_soc__top::commit() {
	bool changed = false;
	if (p_gpio__unit_2e_gpo__reg.commit()) changed = true;
	if (p_gpio__unit_2e_gpio__fsel__reg.commit()) changed = true;
	if (p_gpio__unit_2e_wr__valid.commit()) changed = true;
	if (p_gpo__unit_2e_gpo__reg.commit()) changed = true;
	if (p_gpo__unit_2e_wr__valid__reg.commit()) changed = true;
	if (p_timer__unit_2e_timer__reg.commit()) changed = true;
	if (p_timer__unit_2e_tick__reg.commit()) changed = true;
	if (p_ram__unit_2e_read__word.commit()) changed = true;
	if (p_ram__unit_2e_rd__valid.commit()) changed = true;
	if (p_ram__unit_2e_wr__valid.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__ADDR.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__DATA.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_66_24_104__EN.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__ADDR.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__DATA.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_68_24_105__EN.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__ADDR.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__DATA.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_70_24_106__EN.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__ADDR.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__DATA.commit()) changed = true;
	if (i_flatten_5c_ram__unit_2e__24_memwr_24__5c_mem_24__2e__2f_sys_2f_ram_2e_sv_3a_72_24_107__EN.commit()) changed = true;
	if (p_cpu__unit_2e_state__reg.commit()) changed = true;
	if (p_cpu__unit_2e_alu__result__reg.commit()) changed = true;
	if (p_cpu__unit_2e_store__data__reg.commit()) changed = true;
	if (p_cpu__unit_2e_fetch__unit_2e_pc__reg.commit()) changed = true;
	if (p_cpu__unit_2e_fetch__unit_2e_instr.commit()) changed = true;
	if (p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r0__data.commit()) changed = true;
	if (p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_r1__data.commit()) changed = true;
	if (i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__ADDR.commit()) changed = true;
	if (i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__DATA.commit()) changed = true;
	if (i_flatten_5c_cpu__unit_2e__5c_decode__unit_2e__5c_reg__file__unit_2e__24_memwr_24__5c_regs_24__2e__2f_core_2f_reg__file_2e_sv_3a_26_24_204__EN.commit()) changed = true;
	if (p_cpu__unit_2e_memory__unit_2e_mem__rdata__reg.commit()) changed = true;
	prev_p_clk = p_clk;
	if (memory_p_ram__unit_2e_mem.commit()) changed = true;
	if (memory_p_cpu__unit_2e_decode__unit_2e_reg__file__unit_2e_regs.commit()) changed = true;
	return changed;
}

} // namespace cxxrtl_design

