package model

// func Generate(pattern string, params Params) (*Model, error) {
// 	klines, files, err := kline.LoadKlines(pattern)
// 	if err != nil {
// 		slog.Error("Failed to load klines", "error", err)
// 		return nil, err
// 	}

// 	slog.Info(fmt.Sprintf("%d klines loaded.", len(klines)))

// 	probs := calculateProbs(params, klines)

// 	slog.Info("Verifying probes...")

// 	if err := verifyProbes(params, probs); err != nil {
// 		return nil, err
// 	}

// 	model := &Model{
// 		Probs:   probs,
// 		Params:  params,
// 		Files:   files,
// 		Created: time.Now(),
// 	}

// 	slog.Info("Done.")

// 	return model, nil
// }
