#[unsafe(no_mangle)]
pub extern "C" fn estimate_rank(score: u32, total_users: u32) -> u32 {
    const C0: f32 = -23.72406915;
    const C1: f32 = 7.42490985;
    const C2: f32 = -0.44577973;
    
    const A: f32 = 7.854360;
    const B: f32 = -8.60431e-05;

    const A_LOW: f32 = 12.899220;
    const B_LOW: f32 = -0.00034972;

    const BLEND_LO: f32 = 20000.0;
    const BLEND_HI: f32 = 25000.0;

    if score == 0 {
        return total_users;
    }

    let score_f = score as f32;

    let r: f32 = if score_f >= BLEND_HI {
        let x = score_f.ln();
        (C0 + C1 * x + C2 * x * x).exp()
    } else if score_f <= BLEND_LO {
        (A_LOW + B_LOW * score_f).exp()
    } else {
        let mut t = (score_f - BLEND_LO) / (BLEND_HI - BLEND_LO);
        t = t * t * (3.0 - 2.0 * t);

        let x = score_f.ln();
        let lo = A + B * score_f;
        let hi = C0 + C1 * x + C2 * x * x;

        (lo + t * (hi - lo)).exp()
    };

    r.max(1.0).min(total_users as f32) as u32
}