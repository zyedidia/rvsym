#include "or1200.cxx"
#include "rvsym.h"
#include "constraints.h"

int main() {
    cxxrtl_design::p_or1200__cpu cpu;

    unsigned int icpu_dat_i, icpu_ack_i, icpu_rty_i, icpu_err_i, icpu_adr_i, icpu_tag_i;
    /* commented out debug unit signals */
    //  unsigned int du_stall, du_addr, du_dat_du, du_read, du_write, du_dsr;
    //  unsigned int du_dmr1, du_hwbkpt, du_hwbkpt_ls_r, du_flush_pipe;
    unsigned int dcpu_dat_i, dcpu_ack_i, dcpu_rty_i, dcpu_err_i, dcpu_tag_i;
    unsigned int boot_adr_sel_i;
    unsigned int spr_dat_pic, spr_dat_tt, spr_dat_pm, spr_dat_dmmu, spr_dat_immu, spr_dat_du;
    unsigned int mtspr_dc_done, sig_int, sig_tick;

    rvsym_mark_bytes(&icpu_dat_i, sizeof(icpu_dat_i), "icpu_dat_i");
    rvsym_mark_bytes(&icpu_ack_i, sizeof(icpu_ack_i), "icpu_ack_i");
    rvsym_mark_bytes(&icpu_rty_i, sizeof(icpu_rty_i), "icpu_rty_i");
    rvsym_mark_bytes(&icpu_err_i, sizeof(icpu_err_i), "icpu_err_i");
    rvsym_mark_bytes(&icpu_adr_i, sizeof(icpu_adr_i), "icpu_adr_i");
    rvsym_mark_bytes(&icpu_tag_i, sizeof(icpu_tag_i), "icpu_tag_i");

    /* commented out debug unit signals */
    //  rvsym_mark_bytes(&du_stall, sizeof(du_stall), "du_stall");
    //  rvsym_mark_bytes(&du_addr, sizeof(du_addr), "du_addr");
    //  rvsym_mark_bytes(&du_dat_du, sizeof(du_dat_du), "du_dat_du");
    //  rvsym_mark_bytes(&du_read, sizeof(du_read), "du_read");
    //  rvsym_mark_bytes(&du_write, sizeof(du_write), "du_write");
    //  rvsym_mark_bytes(&du_dsr, sizeof(du_dsr), "du_dsr");

    //  rvsym_mark_bytes(&du_dmr1, sizeof(du_dmr1), "du_dmr1");
    //  rvsym_mark_bytes(&du_hwbkpt, sizeof(du_hwbkpt), "du_hwbkpt");
    //  rvsym_mark_bytes(&du_hwbkpt_ls_r, sizeof(du_hwbkpt_ls_r), "du_hwbkpt_ls_r");
    //  rvsym_mark_bytes(&du_flush_pipe, sizeof(du_flush_pipe), "du_flush_pipe");

    rvsym_mark_bytes(&dcpu_dat_i, sizeof(dcpu_dat_i), "dcpu_dat_i");
    rvsym_mark_bytes(&dcpu_ack_i, sizeof(dcpu_ack_i), "dcpu_ack_i");
    rvsym_mark_bytes(&dcpu_rty_i, sizeof(dcpu_rty_i), "dcpu_rty_i");
    rvsym_mark_bytes(&dcpu_err_i, sizeof(dcpu_err_i), "dcpu_err_i");
    rvsym_mark_bytes(&dcpu_tag_i, sizeof(dcpu_tag_i), "dcpu_tag_i");

    rvsym_mark_bytes(&boot_adr_sel_i, sizeof(boot_adr_sel_i), "boot_adr_sel_i");

    rvsym_mark_bytes(&spr_dat_pic, sizeof(spr_dat_pic), "spr_dat_pic");
    rvsym_mark_bytes(&spr_dat_tt, sizeof(spr_dat_tt), "spr_dat_tt");
    rvsym_mark_bytes(&spr_dat_pm, sizeof(spr_dat_pm), "spr_dat_pm");
    rvsym_mark_bytes(&spr_dat_dmmu, sizeof(spr_dat_dmmu), "spr_dat_dmmu");
    rvsym_mark_bytes(&spr_dat_immu, sizeof(spr_dat_immu), "spr_dat_immu");
    rvsym_mark_bytes(&spr_dat_du, sizeof(spr_dat_du), "spr_dat_du");

    rvsym_mark_bytes(&mtspr_dc_done, sizeof(mtspr_dc_done), "mtspr_dc_done");
    rvsym_mark_bytes(&sig_int, sizeof(sig_int), "sig_int");
    rvsym_mark_bytes(&sig_tick, sizeof(sig_tick), "sig_tick");

    or1k_constraints(icpu_dat_i);

    cpu.p_icpu__dat__i.set<unsigned>(icpu_dat_i);
    cpu.p_icpu__ack__i.set<unsigned>(icpu_ack_i);
    cpu.p_icpu__rty__i.set<unsigned>(icpu_rty_i);
    cpu.p_icpu__err__i.set<unsigned>(icpu_err_i);
    cpu.p_icpu__adr__i.set<unsigned>(icpu_adr_i);
    cpu.p_icpu__tag__i.set<unsigned>(icpu_tag_i);

    /* commented out debug unit signals */
    //  cpu.du_stall = du_stall;
    //  cpu.du_addr = du_addr;
    //  cpu.du_dat_du = du_dat_du;
    //  cpu.du_read = du_read;
    //  cpu.du_write = du_write;
    //  cpu.du_dsr = du_dsr;
    //  cpu.du_dmr1 = du_dmr1;
    //  cpu.du_hwbkpt = du_hwbkpt;
    //  cpu.du_hwbkpt_ls_r = du_hwbkpt_ls_r;
    //  cpu.du_flush_pipe = du_flush_pipe;

    // cpu.du_stall = 0;
    // cpu.du_addr = 0;
    // cpu.du_dat_du = 0;
    // cpu.du_read = 0;
    // cpu.du_write = 0;
    // cpu.du_dsr = 0;
    // cpu.du_dmr1 = 0;
    // cpu.du_hwbkpt = 0;
    // cpu.du_hwbkpt_ls_r = 0;
    // cpu.du_flush_pipe = 0;

    cpu.p_dcpu__dat__i.set<unsigned>(dcpu_dat_i);
    cpu.p_dcpu__ack__i.set<unsigned>(dcpu_ack_i);
    cpu.p_dcpu__rty__i.set<unsigned>(dcpu_rty_i);
    cpu.p_dcpu__err__i.set<unsigned>(dcpu_err_i);
    cpu.p_dcpu__tag__i.set<unsigned>(dcpu_tag_i);

    cpu.p_spr__dat__pic.set<unsigned>(spr_dat_pic);
    cpu.p_spr__dat__tt.set<unsigned>(spr_dat_tt);
    cpu.p_spr__dat__pm.set<unsigned>(spr_dat_pm);
    cpu.p_spr__dat__dmmu.set<unsigned>(spr_dat_dmmu);
    cpu.p_spr__dat__immu.set<unsigned>(spr_dat_immu);
    cpu.p_spr__dat__du.set<unsigned>(spr_dat_du);

    cpu.p_mtspr__dc__done.set<unsigned>(mtspr_dc_done);
    cpu.p_sig__int.set<unsigned>(sig_int);
    cpu.p_sig__tick.set<unsigned>(sig_tick);

    cpu.p_rst.set<bool>(false);

    cpu.p_clk.set<bool>(false);
    cpu.step();
    cpu.p_clk.set<bool>(true);
    cpu.step();
    cpu.p_clk.set<bool>(false);
    cpu.step();
    cpu.p_clk.set<bool>(true);
    cpu.step();

    cpu.step();

    // if ((cpu.__VlSymsp->TOP__or1200_cpu__or1200_rf.__PVT__rf_we == 1) && 
    //         (cpu.__VlSymsp->TOP__or1200_cpu__or1200_rf.__PVT__rf_addrw == 0) && 
    //         (cpu.__VlSymsp->TOP__or1200_cpu__or1200_rf.__PVT__rf_dataw != 0)) 
    //     rvsym_assert(0);

    if ((cpu.p_or1200__rf_2e_rf__we__allow.get<unsigned>() == 1) && 
            (cpu.p_or1200__ctrl_2e_rf__addrw.get<unsigned>() == 0) && 
            (cpu.p_rf__dataw.get<unsigned>() != 0)) 
        rvsym_assert(0);


    return 0;
}
