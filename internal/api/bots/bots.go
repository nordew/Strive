package bots

import "fmt"

// BotStrategy defines an interface for implementing various bot platforms
type BotStrategy interface {
	Initialize(botToken, webAppURL string) error
	Start()
}

// BotManager manages different bot strategies
type BotManager struct {
	botStrategies map[string]BotStrategy
}

func NewBotManager() *BotManager {
	return &BotManager{
		botStrategies: make(map[string]BotStrategy),
	}
}

func (bm *BotManager) RegisterBot(name string, strategy BotStrategy) {
	bm.botStrategies[name] = strategy
}

func (bm *BotManager) StartBot(name, botToken, webAppURL string) error {
	strategy, exists := bm.botStrategies[name]
	if !exists {
		return fmt.Errorf("bot strategy '%s' not found", name)
	}
	err := strategy.Initialize(botToken, webAppURL)
	if err != nil {
		return err
	}
	strategy.Start()
	return nil
}
